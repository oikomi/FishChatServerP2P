//
// Copyright 2015 Hong Miao (miaohong@miaoghong.org). All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package stun

var StunServersList []string
var dictAttrToVal map[string]string
var dictMsgTypeToVal map[string]string

var dictValToMsgType map[string]string

var dictValToAttr map[string]string

func init() {
	StunServersList = []string{"stun.ekiga.net", "stunserver.org", "stun.ideasip.com", "stun.softjoys.com", "stun.voipbuster.com"}
	
	//dictAttrToVal := make(map[string]string)
	
	dictAttrToVal = map[string]string{
		"MappedAddress": MappedAddress,
        "ResponseAddress": ResponseAddress,
        "ChangeRequest": ChangeRequest,
        "SourceAddress": SourceAddress,
        "ChangedAddress": ChangedAddress,
        "Username": Username,
        "Password": Password,
        "MessageIntegrity": MessageIntegrity,
        "ErrorCode": ErrorCode,
        "UnknownAttribute": UnknownAttribute,
        "ReflectedFrom": ReflectedFrom,
        "XorOnly": XorOnly,
        "XorMappedAddress": XorMappedAddress,
        "ServerName": ServerName,
        "SecondaryAddress": SecondaryAddress,
	}
	
	
	dictMsgTypeToVal = map[string]string{
	    "BindRequestMsg": BindRequestMsg,
	    "BindResponseMsg": BindResponseMsg,
	    "BindErrorResponseMsg": BindErrorResponseMsg,
	    "SharedSecretRequestMsg": SharedSecretRequestMsg,
	    "SharedSecretResponseMsg": SharedSecretResponseMsg,
	    "SharedSecretErrorResponseMsg": SharedSecretErrorResponseMsg,	
	}
}

const (
	Blocked = "Blocked"
	OpenInternet = "Open Internet"
	FullCone = "Full Cone"
	SymmetricUDPFirewall = "Symmetric UDP Firewall"
	RestrictNAT = "Restrict NAT"
	RestrictPortNAT = "Restrict Port NAT"
	SymmetricNAT = "Symmetric NAT"
	ChangedAddressError = "Meet an error, when do Test1 on Changed IP and Port"
)


const (
	//stun attributes
	MappedAddress = "0001"
	ResponseAddress = "0002"
	ChangeRequest = "0003"
	SourceAddress = "0004"
	ChangedAddress = "0005"
	Username = "0006"
	Password = "0007"
	MessageIntegrity = "0008"
	ErrorCode = "0009"
	UnknownAttribute = "000A"
	ReflectedFrom = "000B"
	XorOnly = "0021"
	XorMappedAddress = "8020"
	ServerName = "8022"
	SecondaryAddress = "8050"  // Non standard extention
	
	//types for a stun message
	BindRequestMsg = "0001"
	BindResponseMsg = "0101"
	BindErrorResponseMsg = "0111"
	SharedSecretRequestMsg = "0002"
	SharedSecretResponseMsg = "0102"
	SharedSecretErrorResponseMsg = "0112"
)