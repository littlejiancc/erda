// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: source.proto

package pb

import (
	fmt "fmt"
	math "math"

	_ "github.com/erda-project/erda-proto-go/common/pb"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/protobuf/types/known/structpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *PipelineSource) Validate() error {
	if this.TimeCreated != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TimeCreated); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TimeCreated", err)
		}
	}
	if this.TimeUpdated != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TimeUpdated); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TimeUpdated", err)
		}
	}
	return nil
}
func (this *PipelineSourceCreateRequest) Validate() error {
	return nil
}
func (this *PipelineSourceCreateResponse) Validate() error {
	if this.PipelineSource != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PipelineSource); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PipelineSource", err)
		}
	}
	return nil
}
func (this *PipelineSourceUpdateRequest) Validate() error {
	return nil
}
func (this *PipelineSourceUpdateResponse) Validate() error {
	if this.PipelineSource != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PipelineSource); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PipelineSource", err)
		}
	}
	return nil
}
func (this *PipelineSourceDeleteRequest) Validate() error {
	return nil
}
func (this *PipelineSourceDeleteResponse) Validate() error {
	return nil
}
func (this *PipelineSourceGetRequest) Validate() error {
	return nil
}
func (this *PipelineSourceGetResponse) Validate() error {
	if this.PipelineSource != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PipelineSource); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PipelineSource", err)
		}
	}
	return nil
}
func (this *PipelineSourceListRequest) Validate() error {
	return nil
}
func (this *PipelineSourceListResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
