package repository

import (
	"context"
	"fmt"
	"github.com/Alifarid0011/questionnaire-back-end/constant"
	"github.com/Alifarid0011/questionnaire-back-end/internal/models"
	"github.com/casbin/casbin/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type casbinRepository struct {
	enforcer   *casbin.Enforcer
	collection *mongo.Collection
}

func NewCasbinRepository(enforcer *casbin.Enforcer, db *mongo.Database) CasbinRepository {
	return &casbinRepository{enforcer: enforcer, collection: db.Collection(constant.CasbinRule)}
}

func (r *casbinRepository) Enforce(sub, obj, act, AllowOrDeny string) (bool, error) {
	return r.enforcer.Enforce(sub, obj, act, AllowOrDeny)
}

func (r *casbinRepository) AddPolicy(sub, obj, act, AllowOrDeny string) (bool, error) {
	added, err := r.enforcer.AddPolicy(sub, obj, act, AllowOrDeny)
	if err == nil && added {
		_ = r.enforcer.SavePolicy()
	}
	return added, err
}

func (r *casbinRepository) RemovePolicy(sub, obj, act, AllowOrDeny string) (bool, error) {
	removed, err := r.enforcer.RemovePolicy(sub, obj, act, AllowOrDeny)
	if err == nil && removed {
		_ = r.enforcer.SavePolicy()
	}
	return removed, err
}

func (r *casbinRepository) GetPolicies() ([]models.CasbinPolicy, error) {
	rawPolicies, err := r.enforcer.GetPolicy()
	var policies []models.CasbinPolicy
	for _, p := range rawPolicies {
		if len(p) >= 3 {
			policies = append(policies, models.CasbinPolicy{
				Subject: p[0],
				Action:  p[1],
				Object:  p[2],
			})
		}
	}
	return policies, err
}

func (r *casbinRepository) GetAllRolesFromPolicies(ctx context.Context) ([]string, error) {
	pipeline := mongo.Pipeline{
		{{"$match", bson.D{
			{"ptype", "p"},
			{"v0", bson.D{{"$not", bson.D{{"$regex", "^[0-9a-fA-F]{24}$"}}}}},
		}}},
		{{"$group", bson.D{
			{"_id", "$v0"},
		}}},
	}
	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	roles := make([]string, 0, 1000)
	for cursor.Next(ctx) {
		var doc struct {
			ID string `bson:"_id"`
		}
		if err := cursor.Decode(&doc); err != nil {
			return nil, err
		}
		roles = append(roles, doc.ID)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *casbinRepository) AddGroupingPolicy(child, parent string) (bool, error) {
	added, err := r.enforcer.AddGroupingPolicy(child, parent)
	if err == nil && added {
		_ = r.enforcer.SavePolicy()
	}
	return added, err
}

func (r *casbinRepository) RemoveGroupingPolicy(child, parent string) (bool, error) {
	removed, err := r.enforcer.RemoveGroupingPolicy(child, parent)
	if err != nil {
		log.Printf("Failed to remove grouping policy: %v", err)
	}
	return removed, err
}

func (r *casbinRepository) GetGroupingPolicies() ([][]string, error) {
	return r.enforcer.GetGroupingPolicy()
}

func (r *casbinRepository) GetPermissionsGroupedBySubject() ([]models.SubjectWithPermissions, error) {
	rawPolicies, err := r.enforcer.GetPolicy()
	grouped := make(map[string][]models.Permission)
	for _, p := range rawPolicies {
		if len(p) >= 3 {
			sub := p[0]
			perm := models.Permission{
				Action: p[1],
				Object: p[2],
			}
			grouped[sub] = append(grouped[sub], perm)
		}
	}
	var result []models.SubjectWithPermissions
	for sub, perms := range grouped {
		result = append(result, models.SubjectWithPermissions{
			Subject:     sub,
			Permissions: perms,
		})
	}
	return result, err
}

func (r *casbinRepository) GetRolesForUser(userUID string) ([]string, error) {
	return r.enforcer.GetRolesForUser(userUID)
}

func (r *casbinRepository) GetAllPoliciesRelatedToUser(userUID string) ([]models.PermissionDTO, error, []string) {
	roles, err := r.GetRolesForUser(userUID)
	if err != nil {
		return nil, fmt.Errorf("failed to get roles for user: %w", err), []string{}
	}
	subjects := append(roles, userUID)
	policies, err := r.GetPermissionsBySubjects(subjects)
	return policies, err, roles
}
func (r *casbinRepository) GetPermissionsBySubjects(subjects []string) ([]models.PermissionDTO, error) {
	filter := bson.M{"v0": bson.M{"$in": subjects}, "ptype": "p"}
	cursor, err := r.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	var results []models.PermissionDTO
	for cursor.Next(context.TODO()) {
		var rule struct {
			V0 string `bson:"v0"` // sub
			V1 string `bson:"v1"` // obj
			V2 string `bson:"v2"` // act
			V3 string `bson:"v3"` // eft
		}
		if errDecode := cursor.Decode(&rule); errDecode != nil {
			return nil, errDecode
		}
		results = append(results, models.PermissionDTO{
			Sub: rule.V0,
			Obj: rule.V1,
			Act: rule.V2,
			Eft: rule.V3,
		})
	}
	return results, nil
}
