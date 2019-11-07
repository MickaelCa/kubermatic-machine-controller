package ecs

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ModifyInstanceDeployment invokes the ecs.ModifyInstanceDeployment API synchronously
// api document: https://help.aliyun.com/api/ecs/modifyinstancedeployment.html
func (client *Client) ModifyInstanceDeployment(request *ModifyInstanceDeploymentRequest) (response *ModifyInstanceDeploymentResponse, err error) {
	response = CreateModifyInstanceDeploymentResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyInstanceDeploymentWithChan invokes the ecs.ModifyInstanceDeployment API asynchronously
// api document: https://help.aliyun.com/api/ecs/modifyinstancedeployment.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyInstanceDeploymentWithChan(request *ModifyInstanceDeploymentRequest) (<-chan *ModifyInstanceDeploymentResponse, <-chan error) {
	responseChan := make(chan *ModifyInstanceDeploymentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyInstanceDeployment(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ModifyInstanceDeploymentWithCallback invokes the ecs.ModifyInstanceDeployment API asynchronously
// api document: https://help.aliyun.com/api/ecs/modifyinstancedeployment.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyInstanceDeploymentWithCallback(request *ModifyInstanceDeploymentRequest, callback func(response *ModifyInstanceDeploymentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyInstanceDeploymentResponse
		var err error
		defer close(result)
		response, err = client.ModifyInstanceDeployment(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ModifyInstanceDeploymentRequest is the request struct for api ModifyInstanceDeployment
type ModifyInstanceDeploymentRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DeploymentSetId      string           `position:"Query" name:"DeploymentSetId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	Tenancy              string           `position:"Query" name:"Tenancy"`
	DedicatedHostId      string           `position:"Query" name:"DedicatedHostId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	Force                requests.Boolean `position:"Query" name:"Force"`
	Affinity             string           `position:"Query" name:"Affinity"`
}

// ModifyInstanceDeploymentResponse is the response struct for api ModifyInstanceDeployment
type ModifyInstanceDeploymentResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyInstanceDeploymentRequest creates a request to invoke ModifyInstanceDeployment API
func CreateModifyInstanceDeploymentRequest() (request *ModifyInstanceDeploymentRequest) {
	request = &ModifyInstanceDeploymentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ModifyInstanceDeployment", "ecs", "openAPI")
	return
}

// CreateModifyInstanceDeploymentResponse creates a response to parse from ModifyInstanceDeployment response
func CreateModifyInstanceDeploymentResponse() (response *ModifyInstanceDeploymentResponse) {
	response = &ModifyInstanceDeploymentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
