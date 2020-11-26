package main

import (
	"consts"
	"context"
	"game"
)

// Procedure describe every single step to be executed. It is the smallest unit
// of work in a plan.
type Procedure interface {
	// Name identifies a specific procedure.
	Name() string
	// Do execute the business logic for a specific procedure.
	Do(ctx context.Context) ([]Procedure, error)
}

type GetEnergy struct {
	Amount int
	Source string
}

func (s *GetEnergy) Name() string {
	return "get-energy-" + s.Source + string(s.Amount)
}

func (s *GetEnergy) Do(ctx context.Context) ([]Procedure, error) {
	steps := []Procedure{}
	g := game.Form()

	for _, c := range g.Creeps {
		switch c.Harvest(s.Source) {
		case consts.ERR_NO_BODYPART:
			return steps, nil
		case consts.ERR_NOT_IN_RANGE:
			println("Creep should move damnit!")
			// c.MoveTo(s.Source)
		case consts.ERR_INVALID_TARGET:
			println("Get-Energy given a bad source " + s.Source)
		}
	}
	// tags := []*ec2.Tag{}
	// for k, v := range s.Tags {
	// 	if k == "cluster-name" {
	// 		tags = append(tags, &ec2.Tag{
	// 			Key:   aws.String("Name"),
	// 			Value: aws.String(v),
	// 		})
	// 	}
	// 	tags = append(tags, &ec2.Tag{
	// 		Key:   aws.String(k),
	// 		Value: aws.String(v),
	// 	})
	// }

	// instanceInput := &ec2.RunInstancesInput{
	// 	ImageId:      aws.String("ami-0378588b4ae11ec24"),
	// 	InstanceType: aws.String("t2.micro"),
	// 	//UserData:              &userData,
	// 	MinCount: aws.Int64(1),
	// 	MaxCount: aws.Int64(1),
	// 	SubnetId: s.SubnetID,
	// 	TagSpecifications: []*ec2.TagSpecification{
	// 		{
	// 			ResourceType: aws.String("instance"),
	// 			Tags:         tags,
	// 		},
	// 	},
	// }
	// _, err := s.EC2svc.RunInstances(instanceInput)
	// if err != nil {
	// 	return steps, err
	// }
	return steps, nil
}
