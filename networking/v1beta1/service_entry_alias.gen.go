// Code generated by protoc-gen-alias. DO NOT EDIT.
package v1beta1

import "istio.io/api/networking/v1alpha3"

type ServiceEntry = v1alpha3.ServiceEntry
type ServiceEntry_Location = v1alpha3.ServiceEntry_Location

const ServiceEntry_MESH_EXTERNAL ServiceEntry_Location = v1alpha3.ServiceEntry_MESH_EXTERNAL
const ServiceEntry_MESH_INTERNAL ServiceEntry_Location = v1alpha3.ServiceEntry_MESH_INTERNAL

type ServiceEntry_Resolution = v1alpha3.ServiceEntry_Resolution

const ServiceEntry_NONE ServiceEntry_Resolution = v1alpha3.ServiceEntry_NONE
const ServiceEntry_STATIC ServiceEntry_Resolution = v1alpha3.ServiceEntry_STATIC
const ServiceEntry_DNS ServiceEntry_Resolution = v1alpha3.ServiceEntry_DNS
const ServiceEntry_DNS_ROUND_ROBIN ServiceEntry_Resolution = v1alpha3.ServiceEntry_DNS_ROUND_ROBIN

type ServicePort = v1alpha3.ServicePort