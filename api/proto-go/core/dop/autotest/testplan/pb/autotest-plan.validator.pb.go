// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: autotest-plan.proto

package pb

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/protobuf/types/known/timestamppb"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *TestPlanUpdateByHookRequest) Validate() error {
	if this.Content != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Content); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Content", err)
		}
	}
	return nil
}
func (this *Content) Validate() error {
	for _, item := range this.SubContents {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("SubContents", err)
			}
		}
	}
	return nil
}
func (this *TestPlanUpdateByHookResponse) Validate() error {
	return nil
}