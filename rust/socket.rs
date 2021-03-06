// This file is generated by rust-protobuf 2.24.1. Do not edit
// @generated

// https://github.com/rust-lang/rust-clippy/issues/702
#![allow(unknown_lints)]
#![allow(clippy::all)]

#![allow(unused_attributes)]
#![cfg_attr(rustfmt, rustfmt::skip)]

#![allow(box_pointers)]
#![allow(dead_code)]
#![allow(missing_docs)]
#![allow(non_camel_case_types)]
#![allow(non_snake_case)]
#![allow(non_upper_case_globals)]
#![allow(trivial_casts)]
#![allow(unused_imports)]
#![allow(unused_results)]
//! Generated file from `proto/imp/api/socket/socket.proto`

/// Generated files are compatible only with the same version
/// of protobuf runtime.
// const _PROTOBUF_VERSION_CHECK: () = ::protobuf::VERSION_2_24_1;

#[derive(PartialEq,Clone,Default)]
pub struct SendSocketRequest {
    // message fields
    pub pubkey: ::std::string::String,
    // special fields
    pub unknown_fields: ::protobuf::UnknownFields,
    pub cached_size: ::protobuf::CachedSize,
}

impl<'a> ::std::default::Default for &'a SendSocketRequest {
    fn default() -> &'a SendSocketRequest {
        <SendSocketRequest as ::protobuf::Message>::default_instance()
    }
}

impl SendSocketRequest {
    pub fn new() -> SendSocketRequest {
        ::std::default::Default::default()
    }

    // string pubkey = 1;


    pub fn get_pubkey(&self) -> &str {
        &self.pubkey
    }
    pub fn clear_pubkey(&mut self) {
        self.pubkey.clear();
    }

    // Param is passed by value, moved
    pub fn set_pubkey(&mut self, v: ::std::string::String) {
        self.pubkey = v;
    }

    // Mutable pointer to the field.
    // If field is not initialized, it is initialized with default value first.
    pub fn mut_pubkey(&mut self) -> &mut ::std::string::String {
        &mut self.pubkey
    }

    // Take field
    pub fn take_pubkey(&mut self) -> ::std::string::String {
        ::std::mem::replace(&mut self.pubkey, ::std::string::String::new())
    }
}

impl ::protobuf::Message for SendSocketRequest {
    fn is_initialized(&self) -> bool {
        true
    }

    fn merge_from(&mut self, is: &mut ::protobuf::CodedInputStream<'_>) -> ::protobuf::ProtobufResult<()> {
        while !is.eof()? {
            let (field_number, wire_type) = is.read_tag_unpack()?;
            match field_number {
                1 => {
                    ::protobuf::rt::read_singular_proto3_string_into(wire_type, is, &mut self.pubkey)?;
                },
                _ => {
                    ::protobuf::rt::read_unknown_or_skip_group(field_number, wire_type, is, self.mut_unknown_fields())?;
                },
            };
        }
        ::std::result::Result::Ok(())
    }

    // Compute sizes of nested messages
    #[allow(unused_variables)]
    fn compute_size(&self) -> u32 {
        let mut my_size = 0;
        if !self.pubkey.is_empty() {
            my_size += ::protobuf::rt::string_size(1, &self.pubkey);
        }
        my_size += ::protobuf::rt::unknown_fields_size(self.get_unknown_fields());
        self.cached_size.set(my_size);
        my_size
    }

    fn write_to_with_cached_sizes(&self, os: &mut ::protobuf::CodedOutputStream<'_>) -> ::protobuf::ProtobufResult<()> {
        if !self.pubkey.is_empty() {
            os.write_string(1, &self.pubkey)?;
        }
        os.write_unknown_fields(self.get_unknown_fields())?;
        ::std::result::Result::Ok(())
    }

    fn get_cached_size(&self) -> u32 {
        self.cached_size.get()
    }

    fn get_unknown_fields(&self) -> &::protobuf::UnknownFields {
        &self.unknown_fields
    }

    fn mut_unknown_fields(&mut self) -> &mut ::protobuf::UnknownFields {
        &mut self.unknown_fields
    }

    fn as_any(&self) -> &dyn (::std::any::Any) {
        self as &dyn (::std::any::Any)
    }
    fn as_any_mut(&mut self) -> &mut dyn (::std::any::Any) {
        self as &mut dyn (::std::any::Any)
    }
    fn into_any(self: ::std::boxed::Box<Self>) -> ::std::boxed::Box<dyn (::std::any::Any)> {
        self
    }

    fn descriptor(&self) -> &'static ::protobuf::reflect::MessageDescriptor {
        Self::descriptor_static()
    }

    fn new() -> SendSocketRequest {
        SendSocketRequest::new()
    }

    fn descriptor_static() -> &'static ::protobuf::reflect::MessageDescriptor {
        static descriptor: ::protobuf::rt::LazyV2<::protobuf::reflect::MessageDescriptor> = ::protobuf::rt::LazyV2::INIT;
        descriptor.get(|| {
            let mut fields = ::std::vec::Vec::new();
            fields.push(::protobuf::reflect::accessor::make_simple_field_accessor::<_, ::protobuf::types::ProtobufTypeString>(
                "pubkey",
                |m: &SendSocketRequest| { &m.pubkey },
                |m: &mut SendSocketRequest| { &mut m.pubkey },
            ));
            ::protobuf::reflect::MessageDescriptor::new_pb_name::<SendSocketRequest>(
                "SendSocketRequest",
                fields,
                file_descriptor_proto()
            )
        })
    }

    fn default_instance() -> &'static SendSocketRequest {
        static instance: ::protobuf::rt::LazyV2<SendSocketRequest> = ::protobuf::rt::LazyV2::INIT;
        instance.get(SendSocketRequest::new)
    }
}

impl ::protobuf::Clear for SendSocketRequest {
    fn clear(&mut self) {
        self.pubkey.clear();
        self.unknown_fields.clear();
    }
}

impl ::std::fmt::Debug for SendSocketRequest {
    fn fmt(&self, f: &mut ::std::fmt::Formatter<'_>) -> ::std::fmt::Result {
        ::protobuf::text_format::fmt(self, f)
    }
}

impl ::protobuf::reflect::ProtobufValue for SendSocketRequest {
    fn as_ref(&self) -> ::protobuf::reflect::ReflectValueRef {
        ::protobuf::reflect::ReflectValueRef::Message(self)
    }
}

#[derive(PartialEq,Clone,Default)]
pub struct SendSocketResponse {
    // message fields
    pub id: ::std::string::String,
    // special fields
    pub unknown_fields: ::protobuf::UnknownFields,
    pub cached_size: ::protobuf::CachedSize,
}

impl<'a> ::std::default::Default for &'a SendSocketResponse {
    fn default() -> &'a SendSocketResponse {
        <SendSocketResponse as ::protobuf::Message>::default_instance()
    }
}

impl SendSocketResponse {
    pub fn new() -> SendSocketResponse {
        ::std::default::Default::default()
    }

    // string id = 1;


    pub fn get_id(&self) -> &str {
        &self.id
    }
    pub fn clear_id(&mut self) {
        self.id.clear();
    }

    // Param is passed by value, moved
    pub fn set_id(&mut self, v: ::std::string::String) {
        self.id = v;
    }

    // Mutable pointer to the field.
    // If field is not initialized, it is initialized with default value first.
    pub fn mut_id(&mut self) -> &mut ::std::string::String {
        &mut self.id
    }

    // Take field
    pub fn take_id(&mut self) -> ::std::string::String {
        ::std::mem::replace(&mut self.id, ::std::string::String::new())
    }
}

impl ::protobuf::Message for SendSocketResponse {
    fn is_initialized(&self) -> bool {
        true
    }

    fn merge_from(&mut self, is: &mut ::protobuf::CodedInputStream<'_>) -> ::protobuf::ProtobufResult<()> {
        while !is.eof()? {
            let (field_number, wire_type) = is.read_tag_unpack()?;
            match field_number {
                1 => {
                    ::protobuf::rt::read_singular_proto3_string_into(wire_type, is, &mut self.id)?;
                },
                _ => {
                    ::protobuf::rt::read_unknown_or_skip_group(field_number, wire_type, is, self.mut_unknown_fields())?;
                },
            };
        }
        ::std::result::Result::Ok(())
    }

    // Compute sizes of nested messages
    #[allow(unused_variables)]
    fn compute_size(&self) -> u32 {
        let mut my_size = 0;
        if !self.id.is_empty() {
            my_size += ::protobuf::rt::string_size(1, &self.id);
        }
        my_size += ::protobuf::rt::unknown_fields_size(self.get_unknown_fields());
        self.cached_size.set(my_size);
        my_size
    }

    fn write_to_with_cached_sizes(&self, os: &mut ::protobuf::CodedOutputStream<'_>) -> ::protobuf::ProtobufResult<()> {
        if !self.id.is_empty() {
            os.write_string(1, &self.id)?;
        }
        os.write_unknown_fields(self.get_unknown_fields())?;
        ::std::result::Result::Ok(())
    }

    fn get_cached_size(&self) -> u32 {
        self.cached_size.get()
    }

    fn get_unknown_fields(&self) -> &::protobuf::UnknownFields {
        &self.unknown_fields
    }

    fn mut_unknown_fields(&mut self) -> &mut ::protobuf::UnknownFields {
        &mut self.unknown_fields
    }

    fn as_any(&self) -> &dyn (::std::any::Any) {
        self as &dyn (::std::any::Any)
    }
    fn as_any_mut(&mut self) -> &mut dyn (::std::any::Any) {
        self as &mut dyn (::std::any::Any)
    }
    fn into_any(self: ::std::boxed::Box<Self>) -> ::std::boxed::Box<dyn (::std::any::Any)> {
        self
    }

    fn descriptor(&self) -> &'static ::protobuf::reflect::MessageDescriptor {
        Self::descriptor_static()
    }

    fn new() -> SendSocketResponse {
        SendSocketResponse::new()
    }

    fn descriptor_static() -> &'static ::protobuf::reflect::MessageDescriptor {
        static descriptor: ::protobuf::rt::LazyV2<::protobuf::reflect::MessageDescriptor> = ::protobuf::rt::LazyV2::INIT;
        descriptor.get(|| {
            let mut fields = ::std::vec::Vec::new();
            fields.push(::protobuf::reflect::accessor::make_simple_field_accessor::<_, ::protobuf::types::ProtobufTypeString>(
                "id",
                |m: &SendSocketResponse| { &m.id },
                |m: &mut SendSocketResponse| { &mut m.id },
            ));
            ::protobuf::reflect::MessageDescriptor::new_pb_name::<SendSocketResponse>(
                "SendSocketResponse",
                fields,
                file_descriptor_proto()
            )
        })
    }

    fn default_instance() -> &'static SendSocketResponse {
        static instance: ::protobuf::rt::LazyV2<SendSocketResponse> = ::protobuf::rt::LazyV2::INIT;
        instance.get(SendSocketResponse::new)
    }
}

impl ::protobuf::Clear for SendSocketResponse {
    fn clear(&mut self) {
        self.id.clear();
        self.unknown_fields.clear();
    }
}

impl ::std::fmt::Debug for SendSocketResponse {
    fn fmt(&self, f: &mut ::std::fmt::Formatter<'_>) -> ::std::fmt::Result {
        ::protobuf::text_format::fmt(self, f)
    }
}

impl ::protobuf::reflect::ProtobufValue for SendSocketResponse {
    fn as_ref(&self) -> ::protobuf::reflect::ReflectValueRef {
        ::protobuf::reflect::ReflectValueRef::Message(self)
    }
}

#[derive(PartialEq,Clone,Default)]
pub struct StartSocketRequest {
    // special fields
    pub unknown_fields: ::protobuf::UnknownFields,
    pub cached_size: ::protobuf::CachedSize,
}

impl<'a> ::std::default::Default for &'a StartSocketRequest {
    fn default() -> &'a StartSocketRequest {
        <StartSocketRequest as ::protobuf::Message>::default_instance()
    }
}

impl StartSocketRequest {
    pub fn new() -> StartSocketRequest {
        ::std::default::Default::default()
    }
}

impl ::protobuf::Message for StartSocketRequest {
    fn is_initialized(&self) -> bool {
        true
    }

    fn merge_from(&mut self, is: &mut ::protobuf::CodedInputStream<'_>) -> ::protobuf::ProtobufResult<()> {
        while !is.eof()? {
            let (field_number, wire_type) = is.read_tag_unpack()?;
            match field_number {
                _ => {
                    ::protobuf::rt::read_unknown_or_skip_group(field_number, wire_type, is, self.mut_unknown_fields())?;
                },
            };
        }
        ::std::result::Result::Ok(())
    }

    // Compute sizes of nested messages
    #[allow(unused_variables)]
    fn compute_size(&self) -> u32 {
        let mut my_size = 0;
        my_size += ::protobuf::rt::unknown_fields_size(self.get_unknown_fields());
        self.cached_size.set(my_size);
        my_size
    }

    fn write_to_with_cached_sizes(&self, os: &mut ::protobuf::CodedOutputStream<'_>) -> ::protobuf::ProtobufResult<()> {
        os.write_unknown_fields(self.get_unknown_fields())?;
        ::std::result::Result::Ok(())
    }

    fn get_cached_size(&self) -> u32 {
        self.cached_size.get()
    }

    fn get_unknown_fields(&self) -> &::protobuf::UnknownFields {
        &self.unknown_fields
    }

    fn mut_unknown_fields(&mut self) -> &mut ::protobuf::UnknownFields {
        &mut self.unknown_fields
    }

    fn as_any(&self) -> &dyn (::std::any::Any) {
        self as &dyn (::std::any::Any)
    }
    fn as_any_mut(&mut self) -> &mut dyn (::std::any::Any) {
        self as &mut dyn (::std::any::Any)
    }
    fn into_any(self: ::std::boxed::Box<Self>) -> ::std::boxed::Box<dyn (::std::any::Any)> {
        self
    }

    fn descriptor(&self) -> &'static ::protobuf::reflect::MessageDescriptor {
        Self::descriptor_static()
    }

    fn new() -> StartSocketRequest {
        StartSocketRequest::new()
    }

    fn descriptor_static() -> &'static ::protobuf::reflect::MessageDescriptor {
        static descriptor: ::protobuf::rt::LazyV2<::protobuf::reflect::MessageDescriptor> = ::protobuf::rt::LazyV2::INIT;
        descriptor.get(|| {
            let fields = ::std::vec::Vec::new();
            ::protobuf::reflect::MessageDescriptor::new_pb_name::<StartSocketRequest>(
                "StartSocketRequest",
                fields,
                file_descriptor_proto()
            )
        })
    }

    fn default_instance() -> &'static StartSocketRequest {
        static instance: ::protobuf::rt::LazyV2<StartSocketRequest> = ::protobuf::rt::LazyV2::INIT;
        instance.get(StartSocketRequest::new)
    }
}

impl ::protobuf::Clear for StartSocketRequest {
    fn clear(&mut self) {
        self.unknown_fields.clear();
    }
}

impl ::std::fmt::Debug for StartSocketRequest {
    fn fmt(&self, f: &mut ::std::fmt::Formatter<'_>) -> ::std::fmt::Result {
        ::protobuf::text_format::fmt(self, f)
    }
}

impl ::protobuf::reflect::ProtobufValue for StartSocketRequest {
    fn as_ref(&self) -> ::protobuf::reflect::ReflectValueRef {
        ::protobuf::reflect::ReflectValueRef::Message(self)
    }
}

#[derive(PartialEq,Clone,Default)]
pub struct StartSocketResponse {
    // message fields
    pub protocol: ::std::string::String,
    pub ip: ::std::string::String,
    pub port1: ::std::string::String,
    pub port2: ::std::string::String,
    // special fields
    pub unknown_fields: ::protobuf::UnknownFields,
    pub cached_size: ::protobuf::CachedSize,
}

impl<'a> ::std::default::Default for &'a StartSocketResponse {
    fn default() -> &'a StartSocketResponse {
        <StartSocketResponse as ::protobuf::Message>::default_instance()
    }
}

impl StartSocketResponse {
    pub fn new() -> StartSocketResponse {
        ::std::default::Default::default()
    }

    // string protocol = 1;


    pub fn get_protocol(&self) -> &str {
        &self.protocol
    }
    pub fn clear_protocol(&mut self) {
        self.protocol.clear();
    }

    // Param is passed by value, moved
    pub fn set_protocol(&mut self, v: ::std::string::String) {
        self.protocol = v;
    }

    // Mutable pointer to the field.
    // If field is not initialized, it is initialized with default value first.
    pub fn mut_protocol(&mut self) -> &mut ::std::string::String {
        &mut self.protocol
    }

    // Take field
    pub fn take_protocol(&mut self) -> ::std::string::String {
        ::std::mem::replace(&mut self.protocol, ::std::string::String::new())
    }

    // string ip = 2;


    pub fn get_ip(&self) -> &str {
        &self.ip
    }
    pub fn clear_ip(&mut self) {
        self.ip.clear();
    }

    // Param is passed by value, moved
    pub fn set_ip(&mut self, v: ::std::string::String) {
        self.ip = v;
    }

    // Mutable pointer to the field.
    // If field is not initialized, it is initialized with default value first.
    pub fn mut_ip(&mut self) -> &mut ::std::string::String {
        &mut self.ip
    }

    // Take field
    pub fn take_ip(&mut self) -> ::std::string::String {
        ::std::mem::replace(&mut self.ip, ::std::string::String::new())
    }

    // string port1 = 3;


    pub fn get_port1(&self) -> &str {
        &self.port1
    }
    pub fn clear_port1(&mut self) {
        self.port1.clear();
    }

    // Param is passed by value, moved
    pub fn set_port1(&mut self, v: ::std::string::String) {
        self.port1 = v;
    }

    // Mutable pointer to the field.
    // If field is not initialized, it is initialized with default value first.
    pub fn mut_port1(&mut self) -> &mut ::std::string::String {
        &mut self.port1
    }

    // Take field
    pub fn take_port1(&mut self) -> ::std::string::String {
        ::std::mem::replace(&mut self.port1, ::std::string::String::new())
    }

    // string port2 = 4;


    pub fn get_port2(&self) -> &str {
        &self.port2
    }
    pub fn clear_port2(&mut self) {
        self.port2.clear();
    }

    // Param is passed by value, moved
    pub fn set_port2(&mut self, v: ::std::string::String) {
        self.port2 = v;
    }

    // Mutable pointer to the field.
    // If field is not initialized, it is initialized with default value first.
    pub fn mut_port2(&mut self) -> &mut ::std::string::String {
        &mut self.port2
    }

    // Take field
    pub fn take_port2(&mut self) -> ::std::string::String {
        ::std::mem::replace(&mut self.port2, ::std::string::String::new())
    }
}

impl ::protobuf::Message for StartSocketResponse {
    fn is_initialized(&self) -> bool {
        true
    }

    fn merge_from(&mut self, is: &mut ::protobuf::CodedInputStream<'_>) -> ::protobuf::ProtobufResult<()> {
        while !is.eof()? {
            let (field_number, wire_type) = is.read_tag_unpack()?;
            match field_number {
                1 => {
                    ::protobuf::rt::read_singular_proto3_string_into(wire_type, is, &mut self.protocol)?;
                },
                2 => {
                    ::protobuf::rt::read_singular_proto3_string_into(wire_type, is, &mut self.ip)?;
                },
                3 => {
                    ::protobuf::rt::read_singular_proto3_string_into(wire_type, is, &mut self.port1)?;
                },
                4 => {
                    ::protobuf::rt::read_singular_proto3_string_into(wire_type, is, &mut self.port2)?;
                },
                _ => {
                    ::protobuf::rt::read_unknown_or_skip_group(field_number, wire_type, is, self.mut_unknown_fields())?;
                },
            };
        }
        ::std::result::Result::Ok(())
    }

    // Compute sizes of nested messages
    #[allow(unused_variables)]
    fn compute_size(&self) -> u32 {
        let mut my_size = 0;
        if !self.protocol.is_empty() {
            my_size += ::protobuf::rt::string_size(1, &self.protocol);
        }
        if !self.ip.is_empty() {
            my_size += ::protobuf::rt::string_size(2, &self.ip);
        }
        if !self.port1.is_empty() {
            my_size += ::protobuf::rt::string_size(3, &self.port1);
        }
        if !self.port2.is_empty() {
            my_size += ::protobuf::rt::string_size(4, &self.port2);
        }
        my_size += ::protobuf::rt::unknown_fields_size(self.get_unknown_fields());
        self.cached_size.set(my_size);
        my_size
    }

    fn write_to_with_cached_sizes(&self, os: &mut ::protobuf::CodedOutputStream<'_>) -> ::protobuf::ProtobufResult<()> {
        if !self.protocol.is_empty() {
            os.write_string(1, &self.protocol)?;
        }
        if !self.ip.is_empty() {
            os.write_string(2, &self.ip)?;
        }
        if !self.port1.is_empty() {
            os.write_string(3, &self.port1)?;
        }
        if !self.port2.is_empty() {
            os.write_string(4, &self.port2)?;
        }
        os.write_unknown_fields(self.get_unknown_fields())?;
        ::std::result::Result::Ok(())
    }

    fn get_cached_size(&self) -> u32 {
        self.cached_size.get()
    }

    fn get_unknown_fields(&self) -> &::protobuf::UnknownFields {
        &self.unknown_fields
    }

    fn mut_unknown_fields(&mut self) -> &mut ::protobuf::UnknownFields {
        &mut self.unknown_fields
    }

    fn as_any(&self) -> &dyn (::std::any::Any) {
        self as &dyn (::std::any::Any)
    }
    fn as_any_mut(&mut self) -> &mut dyn (::std::any::Any) {
        self as &mut dyn (::std::any::Any)
    }
    fn into_any(self: ::std::boxed::Box<Self>) -> ::std::boxed::Box<dyn (::std::any::Any)> {
        self
    }

    fn descriptor(&self) -> &'static ::protobuf::reflect::MessageDescriptor {
        Self::descriptor_static()
    }

    fn new() -> StartSocketResponse {
        StartSocketResponse::new()
    }

    fn descriptor_static() -> &'static ::protobuf::reflect::MessageDescriptor {
        static descriptor: ::protobuf::rt::LazyV2<::protobuf::reflect::MessageDescriptor> = ::protobuf::rt::LazyV2::INIT;
        descriptor.get(|| {
            let mut fields = ::std::vec::Vec::new();
            fields.push(::protobuf::reflect::accessor::make_simple_field_accessor::<_, ::protobuf::types::ProtobufTypeString>(
                "protocol",
                |m: &StartSocketResponse| { &m.protocol },
                |m: &mut StartSocketResponse| { &mut m.protocol },
            ));
            fields.push(::protobuf::reflect::accessor::make_simple_field_accessor::<_, ::protobuf::types::ProtobufTypeString>(
                "ip",
                |m: &StartSocketResponse| { &m.ip },
                |m: &mut StartSocketResponse| { &mut m.ip },
            ));
            fields.push(::protobuf::reflect::accessor::make_simple_field_accessor::<_, ::protobuf::types::ProtobufTypeString>(
                "port1",
                |m: &StartSocketResponse| { &m.port1 },
                |m: &mut StartSocketResponse| { &mut m.port1 },
            ));
            fields.push(::protobuf::reflect::accessor::make_simple_field_accessor::<_, ::protobuf::types::ProtobufTypeString>(
                "port2",
                |m: &StartSocketResponse| { &m.port2 },
                |m: &mut StartSocketResponse| { &mut m.port2 },
            ));
            ::protobuf::reflect::MessageDescriptor::new_pb_name::<StartSocketResponse>(
                "StartSocketResponse",
                fields,
                file_descriptor_proto()
            )
        })
    }

    fn default_instance() -> &'static StartSocketResponse {
        static instance: ::protobuf::rt::LazyV2<StartSocketResponse> = ::protobuf::rt::LazyV2::INIT;
        instance.get(StartSocketResponse::new)
    }
}

impl ::protobuf::Clear for StartSocketResponse {
    fn clear(&mut self) {
        self.protocol.clear();
        self.ip.clear();
        self.port1.clear();
        self.port2.clear();
        self.unknown_fields.clear();
    }
}

impl ::std::fmt::Debug for StartSocketResponse {
    fn fmt(&self, f: &mut ::std::fmt::Formatter<'_>) -> ::std::fmt::Result {
        ::protobuf::text_format::fmt(self, f)
    }
}

impl ::protobuf::reflect::ProtobufValue for StartSocketResponse {
    fn as_ref(&self) -> ::protobuf::reflect::ReflectValueRef {
        ::protobuf::reflect::ReflectValueRef::Message(self)
    }
}

#[derive(PartialEq,Clone,Default)]
pub struct StopSocketRequest {
    // special fields
    pub unknown_fields: ::protobuf::UnknownFields,
    pub cached_size: ::protobuf::CachedSize,
}

impl<'a> ::std::default::Default for &'a StopSocketRequest {
    fn default() -> &'a StopSocketRequest {
        <StopSocketRequest as ::protobuf::Message>::default_instance()
    }
}

impl StopSocketRequest {
    pub fn new() -> StopSocketRequest {
        ::std::default::Default::default()
    }
}

impl ::protobuf::Message for StopSocketRequest {
    fn is_initialized(&self) -> bool {
        true
    }

    fn merge_from(&mut self, is: &mut ::protobuf::CodedInputStream<'_>) -> ::protobuf::ProtobufResult<()> {
        while !is.eof()? {
            let (field_number, wire_type) = is.read_tag_unpack()?;
            match field_number {
                _ => {
                    ::protobuf::rt::read_unknown_or_skip_group(field_number, wire_type, is, self.mut_unknown_fields())?;
                },
            };
        }
        ::std::result::Result::Ok(())
    }

    // Compute sizes of nested messages
    #[allow(unused_variables)]
    fn compute_size(&self) -> u32 {
        let mut my_size = 0;
        my_size += ::protobuf::rt::unknown_fields_size(self.get_unknown_fields());
        self.cached_size.set(my_size);
        my_size
    }

    fn write_to_with_cached_sizes(&self, os: &mut ::protobuf::CodedOutputStream<'_>) -> ::protobuf::ProtobufResult<()> {
        os.write_unknown_fields(self.get_unknown_fields())?;
        ::std::result::Result::Ok(())
    }

    fn get_cached_size(&self) -> u32 {
        self.cached_size.get()
    }

    fn get_unknown_fields(&self) -> &::protobuf::UnknownFields {
        &self.unknown_fields
    }

    fn mut_unknown_fields(&mut self) -> &mut ::protobuf::UnknownFields {
        &mut self.unknown_fields
    }

    fn as_any(&self) -> &dyn (::std::any::Any) {
        self as &dyn (::std::any::Any)
    }
    fn as_any_mut(&mut self) -> &mut dyn (::std::any::Any) {
        self as &mut dyn (::std::any::Any)
    }
    fn into_any(self: ::std::boxed::Box<Self>) -> ::std::boxed::Box<dyn (::std::any::Any)> {
        self
    }

    fn descriptor(&self) -> &'static ::protobuf::reflect::MessageDescriptor {
        Self::descriptor_static()
    }

    fn new() -> StopSocketRequest {
        StopSocketRequest::new()
    }

    fn descriptor_static() -> &'static ::protobuf::reflect::MessageDescriptor {
        static descriptor: ::protobuf::rt::LazyV2<::protobuf::reflect::MessageDescriptor> = ::protobuf::rt::LazyV2::INIT;
        descriptor.get(|| {
            let fields = ::std::vec::Vec::new();
            ::protobuf::reflect::MessageDescriptor::new_pb_name::<StopSocketRequest>(
                "StopSocketRequest",
                fields,
                file_descriptor_proto()
            )
        })
    }

    fn default_instance() -> &'static StopSocketRequest {
        static instance: ::protobuf::rt::LazyV2<StopSocketRequest> = ::protobuf::rt::LazyV2::INIT;
        instance.get(StopSocketRequest::new)
    }
}

impl ::protobuf::Clear for StopSocketRequest {
    fn clear(&mut self) {
        self.unknown_fields.clear();
    }
}

impl ::std::fmt::Debug for StopSocketRequest {
    fn fmt(&self, f: &mut ::std::fmt::Formatter<'_>) -> ::std::fmt::Result {
        ::protobuf::text_format::fmt(self, f)
    }
}

impl ::protobuf::reflect::ProtobufValue for StopSocketRequest {
    fn as_ref(&self) -> ::protobuf::reflect::ReflectValueRef {
        ::protobuf::reflect::ReflectValueRef::Message(self)
    }
}

#[derive(PartialEq,Clone,Default)]
pub struct StopSocketResponse {
    // special fields
    pub unknown_fields: ::protobuf::UnknownFields,
    pub cached_size: ::protobuf::CachedSize,
}

impl<'a> ::std::default::Default for &'a StopSocketResponse {
    fn default() -> &'a StopSocketResponse {
        <StopSocketResponse as ::protobuf::Message>::default_instance()
    }
}

impl StopSocketResponse {
    pub fn new() -> StopSocketResponse {
        ::std::default::Default::default()
    }
}

impl ::protobuf::Message for StopSocketResponse {
    fn is_initialized(&self) -> bool {
        true
    }

    fn merge_from(&mut self, is: &mut ::protobuf::CodedInputStream<'_>) -> ::protobuf::ProtobufResult<()> {
        while !is.eof()? {
            let (field_number, wire_type) = is.read_tag_unpack()?;
            match field_number {
                _ => {
                    ::protobuf::rt::read_unknown_or_skip_group(field_number, wire_type, is, self.mut_unknown_fields())?;
                },
            };
        }
        ::std::result::Result::Ok(())
    }

    // Compute sizes of nested messages
    #[allow(unused_variables)]
    fn compute_size(&self) -> u32 {
        let mut my_size = 0;
        my_size += ::protobuf::rt::unknown_fields_size(self.get_unknown_fields());
        self.cached_size.set(my_size);
        my_size
    }

    fn write_to_with_cached_sizes(&self, os: &mut ::protobuf::CodedOutputStream<'_>) -> ::protobuf::ProtobufResult<()> {
        os.write_unknown_fields(self.get_unknown_fields())?;
        ::std::result::Result::Ok(())
    }

    fn get_cached_size(&self) -> u32 {
        self.cached_size.get()
    }

    fn get_unknown_fields(&self) -> &::protobuf::UnknownFields {
        &self.unknown_fields
    }

    fn mut_unknown_fields(&mut self) -> &mut ::protobuf::UnknownFields {
        &mut self.unknown_fields
    }

    fn as_any(&self) -> &dyn (::std::any::Any) {
        self as &dyn (::std::any::Any)
    }
    fn as_any_mut(&mut self) -> &mut dyn (::std::any::Any) {
        self as &mut dyn (::std::any::Any)
    }
    fn into_any(self: ::std::boxed::Box<Self>) -> ::std::boxed::Box<dyn (::std::any::Any)> {
        self
    }

    fn descriptor(&self) -> &'static ::protobuf::reflect::MessageDescriptor {
        Self::descriptor_static()
    }

    fn new() -> StopSocketResponse {
        StopSocketResponse::new()
    }

    fn descriptor_static() -> &'static ::protobuf::reflect::MessageDescriptor {
        static descriptor: ::protobuf::rt::LazyV2<::protobuf::reflect::MessageDescriptor> = ::protobuf::rt::LazyV2::INIT;
        descriptor.get(|| {
            let fields = ::std::vec::Vec::new();
            ::protobuf::reflect::MessageDescriptor::new_pb_name::<StopSocketResponse>(
                "StopSocketResponse",
                fields,
                file_descriptor_proto()
            )
        })
    }

    fn default_instance() -> &'static StopSocketResponse {
        static instance: ::protobuf::rt::LazyV2<StopSocketResponse> = ::protobuf::rt::LazyV2::INIT;
        instance.get(StopSocketResponse::new)
    }
}

impl ::protobuf::Clear for StopSocketResponse {
    fn clear(&mut self) {
        self.unknown_fields.clear();
    }
}

impl ::std::fmt::Debug for StopSocketResponse {
    fn fmt(&self, f: &mut ::std::fmt::Formatter<'_>) -> ::std::fmt::Result {
        ::protobuf::text_format::fmt(self, f)
    }
}

impl ::protobuf::reflect::ProtobufValue for StopSocketResponse {
    fn as_ref(&self) -> ::protobuf::reflect::ReflectValueRef {
        ::protobuf::reflect::ReflectValueRef::Message(self)
    }
}

static file_descriptor_proto_data: &'static [u8] = b"\
    \n!proto/imp/api/socket/socket.proto\x12\x06socket\x1a\x1cgoogle/api/ann\
    otations.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"+\n\
    \x11SendSocketRequest\x12\x16\n\x06pubkey\x18\x01\x20\x01(\tR\x06pubkey\
    \"$\n\x12SendSocketResponse\x12\x0e\n\x02id\x18\x01\x20\x01(\tR\x02id\"\
    \x14\n\x12StartSocketRequest\"m\n\x13StartSocketResponse\x12\x1a\n\x08pr\
    otocol\x18\x01\x20\x01(\tR\x08protocol\x12\x0e\n\x02ip\x18\x02\x20\x01(\
    \tR\x02ip\x12\x14\n\x05port1\x18\x03\x20\x01(\tR\x05port1\x12\x14\n\x05p\
    ort2\x18\x04\x20\x01(\tR\x05port2\"\x13\n\x11StopSocketRequest\"\x14\n\
    \x12StopSocketResponse2\xb6\x02\n\x06Socket\x12f\n\nSendSocket\x12\x19.s\
    ocket.SendSocketRequest\x1a\x1a.socket.SendSocketResponse\"!\x82\xd3\xe4\
    \x93\x02\x1b\"\x16/v1/socket/sendRequest:\x01*\x12c\n\x0bStartSocket\x12\
    \x1a.socket.StartSocketRequest\x1a\x1b.socket.StartSocketResponse\"\x1b\
    \x82\xd3\xe4\x93\x02\x15\"\x10/v1/socket/start:\x01*\x12_\n\nStopSocket\
    \x12\x19.socket.StopSocketRequest\x1a\x1a.socket.StopSocketResponse\"\
    \x1a\x82\xd3\xe4\x93\x02\x14\"\x0f/v1/socket/stop:\x01*B\xc6\x01Z#github\
    .com/imperviousai/freeimp/gen\x92A\x9d\x01\x12>\n\x0fSocket\x20Services\
    \"&\n\rImpervious\x20AI\x12\x15https://impervious.ai2\x031.0*\x03\x01\
    \x02\x042\x10application/json:\x10application/jsonr2\n\x14Documentation\
    \x20on\x20IMP\x12\x1ahttps://docs.impervious.aiJ\xbe\x0e\n\x06\x12\x04\
    \x01\0j\x01\nN\n\x01\x0c\x12\x03\x01\0\x12\x1aD/\x20Allows\x20for\x20p2p\
    \x20sockets\x20to\x20be\x20established\x20between\x20Impervious\x20nodes\
    \n\n\x08\n\x01\x02\x12\x03\x03\0\x0f\n\x08\n\x01\x08\x12\x03\x05\0:\n\t\
    \n\x02\x08\x0b\x12\x03\x05\0:\n\t\n\x02\x03\0\x12\x03\x07\0&\n\t\n\x02\
    \x03\x01\x12\x03\x08\08\n\t\n\x01\x08\x12\x04\n\0\x1c\x02\n\x0b\n\x03\
    \x08\x92\x08\x12\x04\n\0\x1c\x02\na\n\x02\x06\0\x12\x04!\0@\x01\x1aU*\n\
    \x20Socket\x20service\x20allows\x20for\x20p2p\x20sockets\x20to\x20be\x20\
    established\x20between\x20Impervious\x20nodes.\n\n\n\n\x03\x06\0\x01\x12\
    \x03!\x08\x0e\nD\n\x04\x06\0\x02\0\x12\x04&\x08+\t\x1a6*\n\x20SendSocket\
    \x20sends\x20a\x20socket\x20request\x20to\x20another\x20node.\n\n\x0c\n\
    \x05\x06\0\x02\0\x01\x12\x03&\x0c\x16\n\x0c\n\x05\x06\0\x02\0\x02\x12\
    \x03&\x17(\n\x0c\n\x05\x06\0\x02\0\x03\x12\x03&3E\n\r\n\x05\x06\0\x02\0\
    \x04\x12\x04'\x10*\x12\n\x11\n\t\x06\0\x02\0\x04\xb0\xca\xbc\"\x12\x04'\
    \x10*\x12\nH\n\x04\x06\0\x02\x01\x12\x040\x085\t\x1a:*\n\x20StartSocket\
    \x20starts\x20the\x20socket\x20on\x20your\x20Impervious\x20node.\n\n\x0c\
    \n\x05\x06\0\x02\x01\x01\x12\x030\x0c\x17\n\x0c\n\x05\x06\0\x02\x01\x02\
    \x12\x030\x18*\n\x0c\n\x05\x06\0\x02\x01\x03\x12\x0305H\n\r\n\x05\x06\0\
    \x02\x01\x04\x12\x041\x104\x12\n\x11\n\t\x06\0\x02\x01\x04\xb0\xca\xbc\"\
    \x12\x041\x104\x12\nF\n\x04\x06\0\x02\x02\x12\x04:\x08?\t\x1a8*\n\x20Sto\
    pSocket\x20stops\x20the\x20socket\x20on\x20your\x20Impervious\x20node.\n\
    \n\x0c\n\x05\x06\0\x02\x02\x01\x12\x03:\x0c\x16\n\x0c\n\x05\x06\0\x02\
    \x02\x02\x12\x03:\x17(\n\x0c\n\x05\x06\0\x02\x02\x03\x12\x03:3E\n\r\n\
    \x05\x06\0\x02\x02\x04\x12\x04;\x10>\x12\n\x11\n\t\x06\0\x02\x02\x04\xb0\
    \xca\xbc\"\x12\x04;\x10>\x12\n\\\n\x02\x04\0\x12\x04E\0G\x01\x1aP*\n\x20\
    Represents\x20a\x20request\x20to\x20send\x20socket\x20connection\x20info\
    rmation\x20to\x20a\x20far\x20end\x20node\n\n\n\n\x03\x04\0\x01\x12\x03E\
    \x08\x19\nA\n\x04\x04\0\x02\0\x12\x03F\x08\x1a\"4\x20The\x20public\x20ke\
    y\x20of\x20the\x20far\x20end\x20LND\x20node\x20running\x20IMP\n\n\x0c\n\
    \x05\x04\0\x02\0\x05\x12\x03F\x08\x0e\n\x0c\n\x05\x04\0\x02\0\x01\x12\
    \x03F\x0f\x15\n\x0c\n\x05\x04\0\x02\0\x03\x12\x03F\x18\x19\nQ\n\x02\x04\
    \x01\x12\x04L\0N\x01\x1aE*\n\x20Represents\x20a\x20response\x20back\x20f\
    rom\x20the\x20request\x20to\x20establish\x20a\x20socket\n\n\n\n\x03\x04\
    \x01\x01\x12\x03L\x08\x1a\n\"\n\x04\x04\x01\x02\0\x12\x03M\x08\x16\"\x15\
    \x20returned\x20message\x20ID\n\n\x0c\n\x05\x04\x01\x02\0\x05\x12\x03M\
    \x08\x0e\n\x0c\n\x05\x04\x01\x02\0\x01\x12\x03M\x0f\x11\n\x0c\n\x05\x04\
    \x01\x02\0\x03\x12\x03M\x14\x15\nK\n\x02\x04\x02\x12\x04S\0T\x01\x1a?*\n\
    \x20Represents\x20a\x20request\x20to\x20start\x20a\x20socket\x20on\x20an\
    \x20owned\x20IMP\x20node\n\n\n\n\x03\x04\x02\x01\x12\x03S\x08\x1a\nY\n\
    \x02\x04\x03\x12\x04Y\0^\x01\x1aM*\n\x20Represents\x20a\x20response\x20c\
    ontaining\x20socket\x20information\x20from\x20a\x20started\x20socket\n\n\
    \n\n\x03\x04\x03\x01\x12\x03Y\x08\x1b\n4\n\x04\x04\x03\x02\0\x12\x03Z\
    \x08\x1c\"'\x20the\x20protocol\x20of\x20the\x20socket\x20connection\n\n\
    \x0c\n\x05\x04\x03\x02\0\x05\x12\x03Z\x08\x0e\n\x0c\n\x05\x04\x03\x02\0\
    \x01\x12\x03Z\x0f\x17\n\x0c\n\x05\x04\x03\x02\0\x03\x12\x03Z\x1a\x1b\n3\
    \n\x04\x04\x03\x02\x01\x12\x03[\x08\x16\"&\x20the\x20ip\x20of\x20the\x20\
    IMP\x20node\x20to\x20connect\x20to\n\n\x0c\n\x05\x04\x03\x02\x01\x05\x12\
    \x03[\x08\x0e\n\x0c\n\x05\x04\x03\x02\x01\x01\x12\x03[\x0f\x11\n\x0c\n\
    \x05\x04\x03\x02\x01\x03\x12\x03[\x14\x15\n6\n\x04\x04\x03\x02\x02\x12\
    \x03\\\x08\x19\")\x20the\x20first\x20port\x20of\x20the\x20socket\x20conn\
    ection\n\n\x0c\n\x05\x04\x03\x02\x02\x05\x12\x03\\\x08\x0e\n\x0c\n\x05\
    \x04\x03\x02\x02\x01\x12\x03\\\x0f\x14\n\x0c\n\x05\x04\x03\x02\x02\x03\
    \x12\x03\\\x17\x18\n7\n\x04\x04\x03\x02\x03\x12\x03]\x08\x19\"*\x20the\
    \x20second\x20port\x20of\x20the\x20socket\x20connection\n\n\x0c\n\x05\
    \x04\x03\x02\x03\x05\x12\x03]\x08\x0e\n\x0c\n\x05\x04\x03\x02\x03\x01\
    \x12\x03]\x0f\x14\n\x0c\n\x05\x04\x03\x02\x03\x03\x12\x03]\x17\x18\nJ\n\
    \x02\x04\x04\x12\x04c\0d\x01\x1a>*\n\x20Represents\x20a\x20request\x20to\
    \x20stop\x20a\x20socket\x20on\x20an\x20owned\x20IMP\x20node\n\n\n\n\x03\
    \x04\x04\x01\x12\x03c\x08\x19\n@\n\x02\x04\x05\x12\x04i\0j\x01\x1a4*\n\
    \x20Represents\x20a\x20response\x20back\x20from\x20a\x20stopped\x20socke\
    t\n\n\n\n\x03\x04\x05\x01\x12\x03i\x08\x1ab\x06proto3\
";

static file_descriptor_proto_lazy: ::protobuf::rt::LazyV2<::protobuf::descriptor::FileDescriptorProto> = ::protobuf::rt::LazyV2::INIT;

fn parse_descriptor_proto() -> ::protobuf::descriptor::FileDescriptorProto {
    ::protobuf::Message::parse_from_bytes(file_descriptor_proto_data).unwrap()
}

pub fn file_descriptor_proto() -> &'static ::protobuf::descriptor::FileDescriptorProto {
    file_descriptor_proto_lazy.get(|| {
        parse_descriptor_proto()
    })
}
