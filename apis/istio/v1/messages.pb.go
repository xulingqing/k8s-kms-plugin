//
//Copyright 2018 The Kubernetes Authors.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.10.0
// source: apis/istio/v1/messages.proto

package istio

import (
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type KeyKind int32

const (
	KeyKind_UNKNOWN KeyKind = 0
	KeyKind_AES     KeyKind = 1
	KeyKind_RSA     KeyKind = 2
	KeyKind_ECC     KeyKind = 3
)

// Enum value maps for KeyKind.
var (
	KeyKind_name = map[int32]string{
		0: "UNKNOWN",
		1: "AES",
		2: "RSA",
		3: "ECC",
	}
	KeyKind_value = map[string]int32{
		"UNKNOWN": 0,
		"AES":     1,
		"RSA":     2,
		"ECC":     3,
	}
)

func (x KeyKind) Enum() *KeyKind {
	p := new(KeyKind)
	*p = x
	return p
}

func (x KeyKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KeyKind) Descriptor() protoreflect.EnumDescriptor {
	return file_apis_istio_v1_messages_proto_enumTypes[0].Descriptor()
}

func (KeyKind) Type() protoreflect.EnumType {
	return &file_apis_istio_v1_messages_proto_enumTypes[0]
}

func (x KeyKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KeyKind.Descriptor instead.
func (KeyKind) EnumDescriptor() ([]byte, []int) {
	return file_apis_istio_v1_messages_proto_rawDescGZIP(), []int{0}
}

type Curve int32

const (
	Curve_UNSUPPORTED Curve = 0
	Curve_ED25519     Curve = 1
)

// Enum value maps for Curve.
var (
	Curve_name = map[int32]string{
		0: "UNSUPPORTED",
		1: "ED25519",
	}
	Curve_value = map[string]int32{
		"UNSUPPORTED": 0,
		"ED25519":     1,
	}
)

func (x Curve) Enum() *Curve {
	p := new(Curve)
	*p = x
	return p
}

func (x Curve) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Curve) Descriptor() protoreflect.EnumDescriptor {
	return file_apis_istio_v1_messages_proto_enumTypes[1].Descriptor()
}

func (Curve) Type() protoreflect.EnumType {
	return &file_apis_istio_v1_messages_proto_enumTypes[1]
}

func (x Curve) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Curve.Descriptor instead.
func (Curve) EnumDescriptor() ([]byte, []int) {
	return file_apis_istio_v1_messages_proto_rawDescGZIP(), []int{1}
}

type GenerateKEKRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// optional kid, otherwise will be autogenerated as a UUID.v4 in the response
	KekKid []byte `protobuf:"bytes,1,opt,name=kek_kid,json=kekKid,proto3" json:"kek_kid,omitempty"`
}

func (x *GenerateKEKRequest) Reset() {
	*x = GenerateKEKRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_istio_v1_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateKEKRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateKEKRequest) ProtoMessage() {}

func (x *GenerateKEKRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apis_istio_v1_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateKEKRequest.ProtoReflect.Descriptor instead.
func (*GenerateKEKRequest) Descriptor() ([]byte, []int) {
	return file_apis_istio_v1_messages_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateKEKRequest) GetKekKid() []byte {
	if x != nil {
		return x.KekKid
	}
	return nil
}

type GenerateKEKResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// KEK kid
	KekKid []byte `protobuf:"bytes,1,opt,name=kek_kid,json=kekKid,proto3" json:"kek_kid,omitempty"`
}

func (x *GenerateKEKResponse) Reset() {
	*x = GenerateKEKResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_istio_v1_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateKEKResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateKEKResponse) ProtoMessage() {}

func (x *GenerateKEKResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apis_istio_v1_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateKEKResponse.ProtoReflect.Descriptor instead.
func (*GenerateKEKResponse) Descriptor() ([]byte, []int) {
	return file_apis_istio_v1_messages_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateKEKResponse) GetKekKid() []byte {
	if x != nil {
		return x.KekKid
	}
	return nil
}

type GenerateDEKRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Wrapping/Encrypting KEK ID
	KekKid []byte `protobuf:"bytes,1,opt,name=kek_kid,json=kekKid,proto3" json:"kek_kid,omitempty"`
}

func (x *GenerateDEKRequest) Reset() {
	*x = GenerateDEKRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_istio_v1_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateDEKRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateDEKRequest) ProtoMessage() {}

func (x *GenerateDEKRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apis_istio_v1_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateDEKRequest.ProtoReflect.Descriptor instead.
func (*GenerateDEKRequest) Descriptor() ([]byte, []int) {
	return file_apis_istio_v1_messages_proto_rawDescGZIP(), []int{2}
}

func (x *GenerateDEKRequest) GetKekKid() []byte {
	if x != nil {
		return x.KekKid
	}
	return nil
}

type GenerateDEKResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Encrypted key blob
	EncryptedDekBlob []byte `protobuf:"bytes,1,opt,name=encrypted_dek_blob,json=encryptedDekBlob,proto3" json:"encrypted_dek_blob,omitempty"`
}

func (x *GenerateDEKResponse) Reset() {
	*x = GenerateDEKResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_istio_v1_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateDEKResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateDEKResponse) ProtoMessage() {}

func (x *GenerateDEKResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apis_istio_v1_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateDEKResponse.ProtoReflect.Descriptor instead.
func (*GenerateDEKResponse) Descriptor() ([]byte, []int) {
	return file_apis_istio_v1_messages_proto_rawDescGZIP(), []int{3}
}

func (x *GenerateDEKResponse) GetEncryptedDekBlob() []byte {
	if x != nil {
		return x.EncryptedDekBlob
	}
	return nil
}

type GenerateSKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// key size in bits
	Size int64 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"` // Only supports sizes 2048, 4096 for RSA keys... and
	// What kind of key is it... only Asymmetric kinds will be supported
	Kind KeyKind `protobuf:"varint,2,opt,name=kind,proto3,enum=thalescpl.io.ekms.istio.v1.KeyKind" json:"kind,omitempty"`
	// Curves
	Curve Curve `protobuf:"varint,3,opt,name=curve,proto3,enum=thalescpl.io.ekms.istio.v1.Curve" json:"curve,omitempty"`
	// Parent KID of the KEK
	KekKid []byte `protobuf:"bytes,4,opt,name=kek_kid,json=kekKid,proto3" json:"kek_kid,omitempty"`
	// Encrypted blob of DEK
	EncryptedDekBlob []byte `protobuf:"bytes,5,opt,name=encrypted_dek_blob,json=encryptedDekBlob,proto3" json:"encrypted_dek_blob,omitempty"`
}

func (x *GenerateSKeyRequest) Reset() {
	*x = GenerateSKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_istio_v1_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateSKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateSKeyRequest) ProtoMessage() {}

func (x *GenerateSKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apis_istio_v1_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateSKeyRequest.ProtoReflect.Descriptor instead.
func (*GenerateSKeyRequest) Descriptor() ([]byte, []int) {
	return file_apis_istio_v1_messages_proto_rawDescGZIP(), []int{4}
}

func (x *GenerateSKeyRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *GenerateSKeyRequest) GetKind() KeyKind {
	if x != nil {
		return x.Kind
	}
	return KeyKind_UNKNOWN
}

func (x *GenerateSKeyRequest) GetCurve() Curve {
	if x != nil {
		return x.Curve
	}
	return Curve_UNSUPPORTED
}

func (x *GenerateSKeyRequest) GetKekKid() []byte {
	if x != nil {
		return x.KekKid
	}
	return nil
}

func (x *GenerateSKeyRequest) GetEncryptedDekBlob() []byte {
	if x != nil {
		return x.EncryptedDekBlob
	}
	return nil
}

type GenerateSKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Encrypted blob of SKey encrypted by DEK
	EncryptedSkeyBlob []byte `protobuf:"bytes,1,opt,name=encrypted_skey_blob,json=encryptedSkeyBlob,proto3" json:"encrypted_skey_blob,omitempty"`
}

func (x *GenerateSKeyResponse) Reset() {
	*x = GenerateSKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_istio_v1_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateSKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateSKeyResponse) ProtoMessage() {}

func (x *GenerateSKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apis_istio_v1_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateSKeyResponse.ProtoReflect.Descriptor instead.
func (*GenerateSKeyResponse) Descriptor() ([]byte, []int) {
	return file_apis_istio_v1_messages_proto_rawDescGZIP(), []int{5}
}

func (x *GenerateSKeyResponse) GetEncryptedSkeyBlob() []byte {
	if x != nil {
		return x.EncryptedSkeyBlob
	}
	return nil
}

type LoadSKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// KEK
	KekKid []byte `protobuf:"bytes,1,opt,name=kek_kid,json=kekKid,proto3" json:"kek_kid,omitempty"`
	// Encrypted blob of DEK
	EncryptedDekBlob []byte `protobuf:"bytes,2,opt,name=encrypted_dek_blob,json=encryptedDekBlob,proto3" json:"encrypted_dek_blob,omitempty"`
	// Encrypted blob of SKey
	EncryptedSkeyBlob []byte `protobuf:"bytes,3,opt,name=encrypted_skey_blob,json=encryptedSkeyBlob,proto3" json:"encrypted_skey_blob,omitempty"`
}

func (x *LoadSKeyRequest) Reset() {
	*x = LoadSKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_istio_v1_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadSKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadSKeyRequest) ProtoMessage() {}

func (x *LoadSKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apis_istio_v1_messages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadSKeyRequest.ProtoReflect.Descriptor instead.
func (*LoadSKeyRequest) Descriptor() ([]byte, []int) {
	return file_apis_istio_v1_messages_proto_rawDescGZIP(), []int{6}
}

func (x *LoadSKeyRequest) GetKekKid() []byte {
	if x != nil {
		return x.KekKid
	}
	return nil
}

func (x *LoadSKeyRequest) GetEncryptedDekBlob() []byte {
	if x != nil {
		return x.EncryptedDekBlob
	}
	return nil
}

func (x *LoadSKeyRequest) GetEncryptedSkeyBlob() []byte {
	if x != nil {
		return x.EncryptedSkeyBlob
	}
	return nil
}

type LoadSKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Intermediate CA Signing KeyPair in plaintext in PEM Format
	PlaintextSkey []byte `protobuf:"bytes,1,opt,name=plaintext_skey,json=plaintextSkey,proto3" json:"plaintext_skey,omitempty"`
}

func (x *LoadSKeyResponse) Reset() {
	*x = LoadSKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_istio_v1_messages_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadSKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadSKeyResponse) ProtoMessage() {}

func (x *LoadSKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apis_istio_v1_messages_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadSKeyResponse.ProtoReflect.Descriptor instead.
func (*LoadSKeyResponse) Descriptor() ([]byte, []int) {
	return file_apis_istio_v1_messages_proto_rawDescGZIP(), []int{7}
}

func (x *LoadSKeyResponse) GetPlaintextSkey() []byte {
	if x != nil {
		return x.PlaintextSkey
	}
	return nil
}

// AuthenticatedEncryptRequest
type AuthenticatedEncryptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KekKid []byte `protobuf:"bytes,1,opt,name=kek_kid,json=kekKid,proto3" json:"kek_kid,omitempty"`
		// Generated out of band or via GenerateKEK

	EncryptedDekBlob []byte `protobuf:"bytes,2,opt,name=encrypted_dek_blob,json=encryptedDekBlob,proto3" json:"encrypted_dek_blob,omitempty"` // Encrypted DEK payload wrapped by the KEK
	Plaintext        []byte `protobuf:"bytes,3,opt,name=plaintext,proto3" json:"plaintext,omitempty"`
	Aad              []byte `protobuf:"bytes,4,opt,name=aad,proto3" json:"aad,omitempty"`
}

func (x *AuthenticatedEncryptRequest) Reset() {
	*x = AuthenticatedEncryptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_istio_v1_messages_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticatedEncryptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticatedEncryptRequest) ProtoMessage() {}

func (x *AuthenticatedEncryptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apis_istio_v1_messages_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticatedEncryptRequest.ProtoReflect.Descriptor instead.
func (*AuthenticatedEncryptRequest) Descriptor() ([]byte, []int) {
	return file_apis_istio_v1_messages_proto_rawDescGZIP(), []int{8}
}

func (x *AuthenticatedEncryptRequest) GetKekKid() []byte {
	if x != nil {
		return x.KekKid
	}
	return nil
}

func (x *AuthenticatedEncryptRequest) GetEncryptedDekBlob() []byte {
	if x != nil {
		return x.EncryptedDekBlob
	}
	return nil
}

func (x *AuthenticatedEncryptRequest) GetPlaintext() []byte {
	if x != nil {
		return x.Plaintext
	}
	return nil
}

func (x *AuthenticatedEncryptRequest) GetAad() []byte {
	if x != nil {
		return x.Aad
	}
	return nil
}

// AuthenticatedEncryptResponse
type AuthenticatedEncryptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ciphertext []byte `protobuf:"bytes,1,opt,name=ciphertext,proto3" json:"ciphertext,omitempty"`
}

func (x *AuthenticatedEncryptResponse) Reset() {
	*x = AuthenticatedEncryptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_istio_v1_messages_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticatedEncryptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticatedEncryptResponse) ProtoMessage() {}

func (x *AuthenticatedEncryptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apis_istio_v1_messages_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticatedEncryptResponse.ProtoReflect.Descriptor instead.
func (*AuthenticatedEncryptResponse) Descriptor() ([]byte, []int) {
	return file_apis_istio_v1_messages_proto_rawDescGZIP(), []int{9}
}

func (x *AuthenticatedEncryptResponse) GetCiphertext() []byte {
	if x != nil {
		return x.Ciphertext
	}
	return nil
}

// AuthenticatedDecryptRequest
type AuthenticatedDecryptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KekKid           []byte `protobuf:"bytes,1,opt,name=kek_kid,json=kekKid,proto3" json:"kek_kid,omitempty"`
	EncryptedDekBlob []byte `protobuf:"bytes,2,opt,name=encrypted_dek_blob,json=encryptedDekBlob,proto3" json:"encrypted_dek_blob,omitempty"`
	Ciphertext       []byte `protobuf:"bytes,3,opt,name=ciphertext,proto3" json:"ciphertext,omitempty"`
	Aad              []byte `protobuf:"bytes,4,opt,name=aad,proto3" json:"aad,omitempty"`
}

func (x *AuthenticatedDecryptRequest) Reset() {
	*x = AuthenticatedDecryptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_istio_v1_messages_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticatedDecryptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticatedDecryptRequest) ProtoMessage() {}

func (x *AuthenticatedDecryptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apis_istio_v1_messages_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticatedDecryptRequest.ProtoReflect.Descriptor instead.
func (*AuthenticatedDecryptRequest) Descriptor() ([]byte, []int) {
	return file_apis_istio_v1_messages_proto_rawDescGZIP(), []int{10}
}

func (x *AuthenticatedDecryptRequest) GetKekKid() []byte {
	if x != nil {
		return x.KekKid
	}
	return nil
}

func (x *AuthenticatedDecryptRequest) GetEncryptedDekBlob() []byte {
	if x != nil {
		return x.EncryptedDekBlob
	}
	return nil
}

func (x *AuthenticatedDecryptRequest) GetCiphertext() []byte {
	if x != nil {
		return x.Ciphertext
	}
	return nil
}

func (x *AuthenticatedDecryptRequest) GetAad() []byte {
	if x != nil {
		return x.Aad
	}
	return nil
}

// AuthenticatedDecryptResponse
type AuthenticatedDecryptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plaintext []byte `protobuf:"bytes,1,opt,name=plaintext,proto3" json:"plaintext,omitempty"`
}

func (x *AuthenticatedDecryptResponse) Reset() {
	*x = AuthenticatedDecryptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_istio_v1_messages_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticatedDecryptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticatedDecryptResponse) ProtoMessage() {}

func (x *AuthenticatedDecryptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apis_istio_v1_messages_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticatedDecryptResponse.ProtoReflect.Descriptor instead.
func (*AuthenticatedDecryptResponse) Descriptor() ([]byte, []int) {
	return file_apis_istio_v1_messages_proto_rawDescGZIP(), []int{11}
}

func (x *AuthenticatedDecryptResponse) GetPlaintext() []byte {
	if x != nil {
		return x.Plaintext
	}
	return nil
}

var File_apis_istio_v1_messages_proto protoreflect.FileDescriptor

var file_apis_istio_v1_messages_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a,
	0x74, 0x68, 0x61, 0x6c, 0x65, 0x73, 0x63, 0x70, 0x6c, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6b, 0x6d,
	0x73, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x12, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x4b, 0x45, 0x4b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x6b, 0x65, 0x6b, 0x5f, 0x6b, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x6b, 0x65, 0x6b, 0x4b, 0x69, 0x64, 0x22, 0x2e, 0x0a, 0x13, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x4b, 0x45, 0x4b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x17, 0x0a, 0x07, 0x6b, 0x65, 0x6b, 0x5f, 0x6b, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x6b, 0x65, 0x6b, 0x4b, 0x69, 0x64, 0x22, 0x2d, 0x0a, 0x12, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x44, 0x45, 0x4b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x6b, 0x65, 0x6b, 0x5f, 0x6b, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x06, 0x6b, 0x65, 0x6b, 0x4b, 0x69, 0x64, 0x22, 0x43, 0x0a, 0x13, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x44, 0x45, 0x4b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c,
	0x0a, 0x12, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x65, 0x6b, 0x5f,
	0x62, 0x6c, 0x6f, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x65, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x65, 0x64, 0x44, 0x65, 0x6b, 0x42, 0x6c, 0x6f, 0x62, 0x22, 0xe2, 0x01, 0x0a,
	0x13, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x53, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x74, 0x68, 0x61, 0x6c, 0x65, 0x73, 0x63,
	0x70, 0x6c, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6b, 0x6d, 0x73, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f,
	0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x65, 0x79, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e,
	0x64, 0x12, 0x37, 0x0a, 0x05, 0x63, 0x75, 0x72, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x21, 0x2e, 0x74, 0x68, 0x61, 0x6c, 0x65, 0x73, 0x63, 0x70, 0x6c, 0x2e, 0x69, 0x6f, 0x2e,
	0x65, 0x6b, 0x6d, 0x73, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75,
	0x72, 0x76, 0x65, 0x52, 0x05, 0x63, 0x75, 0x72, 0x76, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6b, 0x65,
	0x6b, 0x5f, 0x6b, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6b, 0x65, 0x6b,
	0x4b, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64,
	0x5f, 0x64, 0x65, 0x6b, 0x5f, 0x62, 0x6c, 0x6f, 0x62, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x10, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x44, 0x65, 0x6b, 0x42, 0x6c, 0x6f,
	0x62, 0x22, 0x46, 0x0a, 0x14, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x53, 0x4b, 0x65,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x65, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x6b, 0x65, 0x79, 0x5f, 0x62, 0x6c, 0x6f, 0x62,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x11, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65,
	0x64, 0x53, 0x6b, 0x65, 0x79, 0x42, 0x6c, 0x6f, 0x62, 0x22, 0x88, 0x01, 0x0a, 0x0f, 0x4c, 0x6f,
	0x61, 0x64, 0x53, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x6b, 0x65, 0x6b, 0x5f, 0x6b, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06,
	0x6b, 0x65, 0x6b, 0x4b, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x65, 0x64, 0x5f, 0x64, 0x65, 0x6b, 0x5f, 0x62, 0x6c, 0x6f, 0x62, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x10, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x44, 0x65, 0x6b,
	0x42, 0x6c, 0x6f, 0x62, 0x12, 0x2e, 0x0a, 0x13, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65,
	0x64, 0x5f, 0x73, 0x6b, 0x65, 0x79, 0x5f, 0x62, 0x6c, 0x6f, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x11, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x53, 0x6b, 0x65, 0x79,
	0x42, 0x6c, 0x6f, 0x62, 0x22, 0x39, 0x0a, 0x10, 0x4c, 0x6f, 0x61, 0x64, 0x53, 0x4b, 0x65, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x6c, 0x61, 0x69,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x73, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x53, 0x6b, 0x65, 0x79, 0x22,
	0x94, 0x01, 0x0a, 0x1b, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x64, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x6b, 0x65, 0x6b, 0x5f, 0x6b, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6b, 0x65, 0x6b, 0x4b, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x65, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x65, 0x6b, 0x5f, 0x62, 0x6c, 0x6f, 0x62, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x44,
	0x65, 0x6b, 0x42, 0x6c, 0x6f, 0x62, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x69, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x03, 0x61, 0x61, 0x64, 0x22, 0x3e, 0x0a, 0x1c, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x63, 0x69, 0x70, 0x68,
	0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x22, 0x96, 0x01, 0x0a, 0x1b, 0x41, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6b, 0x65, 0x6b, 0x5f, 0x6b, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6b, 0x65, 0x6b, 0x4b, 0x69, 0x64, 0x12,
	0x2c, 0x0a, 0x12, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x65, 0x6b,
	0x5f, 0x62, 0x6c, 0x6f, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x65, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x44, 0x65, 0x6b, 0x42, 0x6c, 0x6f, 0x62, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0a, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x61, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x61, 0x61, 0x64, 0x22,
	0x3c, 0x0a, 0x1c, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64,
	0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2a, 0x31, 0x0a,
	0x07, 0x4b, 0x65, 0x79, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x45, 0x53, 0x10, 0x01, 0x12, 0x07,
	0x0a, 0x03, 0x52, 0x53, 0x41, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x45, 0x43, 0x43, 0x10, 0x03,
	0x2a, 0x25, 0x0a, 0x05, 0x43, 0x75, 0x72, 0x76, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53,
	0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x44,
	0x32, 0x35, 0x35, 0x31, 0x39, 0x10, 0x01, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x61, 0x6c, 0x65, 0x73, 0x63, 0x70, 0x6c, 0x2d,
	0x69, 0x6f, 0x2f, 0x6b, 0x38, 0x73, 0x2d, 0x6b, 0x6d, 0x73, 0x2d, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2f, 0x76, 0x31, 0x3b,
	0x69, 0x73, 0x74, 0x69, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apis_istio_v1_messages_proto_rawDescOnce sync.Once
	file_apis_istio_v1_messages_proto_rawDescData = file_apis_istio_v1_messages_proto_rawDesc
)

func file_apis_istio_v1_messages_proto_rawDescGZIP() []byte {
	file_apis_istio_v1_messages_proto_rawDescOnce.Do(func() {
		file_apis_istio_v1_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_apis_istio_v1_messages_proto_rawDescData)
	})
	return file_apis_istio_v1_messages_proto_rawDescData
}

var file_apis_istio_v1_messages_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_apis_istio_v1_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_apis_istio_v1_messages_proto_goTypes = []interface{}{
	(KeyKind)(0),                         // 0: thalescpl.io.ekms.istio.v1.KeyKind
	(Curve)(0),                           // 1: thalescpl.io.ekms.istio.v1.Curve
	(*GenerateKEKRequest)(nil),           // 2: thalescpl.io.ekms.istio.v1.GenerateKEKRequest
	(*GenerateKEKResponse)(nil),          // 3: thalescpl.io.ekms.istio.v1.GenerateKEKResponse
	(*GenerateDEKRequest)(nil),           // 4: thalescpl.io.ekms.istio.v1.GenerateDEKRequest
	(*GenerateDEKResponse)(nil),          // 5: thalescpl.io.ekms.istio.v1.GenerateDEKResponse
	(*GenerateSKeyRequest)(nil),          // 6: thalescpl.io.ekms.istio.v1.GenerateSKeyRequest
	(*GenerateSKeyResponse)(nil),         // 7: thalescpl.io.ekms.istio.v1.GenerateSKeyResponse
	(*LoadSKeyRequest)(nil),              // 8: thalescpl.io.ekms.istio.v1.LoadSKeyRequest
	(*LoadSKeyResponse)(nil),             // 9: thalescpl.io.ekms.istio.v1.LoadSKeyResponse
	(*AuthenticatedEncryptRequest)(nil),  // 10: thalescpl.io.ekms.istio.v1.AuthenticatedEncryptRequest
	(*AuthenticatedEncryptResponse)(nil), // 11: thalescpl.io.ekms.istio.v1.AuthenticatedEncryptResponse
	(*AuthenticatedDecryptRequest)(nil),  // 12: thalescpl.io.ekms.istio.v1.AuthenticatedDecryptRequest
	(*AuthenticatedDecryptResponse)(nil), // 13: thalescpl.io.ekms.istio.v1.AuthenticatedDecryptResponse
}
var file_apis_istio_v1_messages_proto_depIdxs = []int32{
	0, // 0: thalescpl.io.ekms.istio.v1.GenerateSKeyRequest.kind:type_name -> thalescpl.io.ekms.istio.v1.KeyKind
	1, // 1: thalescpl.io.ekms.istio.v1.GenerateSKeyRequest.curve:type_name -> thalescpl.io.ekms.istio.v1.Curve
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_apis_istio_v1_messages_proto_init() }
func file_apis_istio_v1_messages_proto_init() {
	if File_apis_istio_v1_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apis_istio_v1_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateKEKRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apis_istio_v1_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateKEKResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apis_istio_v1_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateDEKRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apis_istio_v1_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateDEKResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apis_istio_v1_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateSKeyRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apis_istio_v1_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateSKeyResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apis_istio_v1_messages_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadSKeyRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apis_istio_v1_messages_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadSKeyResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apis_istio_v1_messages_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticatedEncryptRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apis_istio_v1_messages_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticatedEncryptResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apis_istio_v1_messages_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticatedDecryptRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apis_istio_v1_messages_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticatedDecryptResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apis_istio_v1_messages_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apis_istio_v1_messages_proto_goTypes,
		DependencyIndexes: file_apis_istio_v1_messages_proto_depIdxs,
		EnumInfos:         file_apis_istio_v1_messages_proto_enumTypes,
		MessageInfos:      file_apis_istio_v1_messages_proto_msgTypes,
	}.Build()
	File_apis_istio_v1_messages_proto = out.File
	file_apis_istio_v1_messages_proto_rawDesc = nil
	file_apis_istio_v1_messages_proto_goTypes = nil
	file_apis_istio_v1_messages_proto_depIdxs = nil
}
