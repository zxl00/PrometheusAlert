package edas

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

// App is a nested struct in edas response
type App struct {
	AppId                string `json:"AppId" xml:"AppId"`
	Name                 string `json:"Name" xml:"Name"`
	RegionId             string `json:"RegionId" xml:"RegionId"`
	Description          string `json:"Description" xml:"Description"`
	Owner                string `json:"Owner" xml:"Owner"`
	InstanceCount        int    `json:"InstanceCount" xml:"InstanceCount"`
	RunningInstanceCount int    `json:"RunningInstanceCount" xml:"RunningInstanceCount"`
	Port                 int    `json:"Port" xml:"Port"`
	UserId               string `json:"UserId" xml:"UserId"`
	SlbId                string `json:"SlbId" xml:"SlbId"`
	SlbIp                string `json:"SlbIp" xml:"SlbIp"`
	SlbPort              int    `json:"SlbPort" xml:"SlbPort"`
	ExtSlbId             string `json:"ExtSlbId" xml:"ExtSlbId"`
	ExtSlbIp             string `json:"ExtSlbIp" xml:"ExtSlbIp"`
	ApplicationType      string `json:"ApplicationType" xml:"ApplicationType"`
	ClusterType          int    `json:"ClusterType" xml:"ClusterType"`
	ClusterId            string `json:"ClusterId" xml:"ClusterId"`
	Dockerize            bool   `json:"Dockerize" xml:"Dockerize"`
	Cpu                  int    `json:"Cpu" xml:"Cpu"`
	Memory               int    `json:"Memory" xml:"Memory"`
	HealthCheckUrl       string `json:"HealthCheckUrl" xml:"HealthCheckUrl"`
	BuildPackageId       int    `json:"BuildPackageId" xml:"BuildPackageId"`
	CreateTime           int    `json:"CreateTime" xml:"CreateTime"`
}