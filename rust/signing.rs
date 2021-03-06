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
//! Generated file from `proto/imp/api/signing/signing.proto`

/// Generated files are compatible only with the same version
/// of protobuf runtime.
// const _PROTOBUF_VERSION_CHECK: () = ::protobuf::VERSION_2_24_1;

#[derive(PartialEq,Clone,Default)]
pub struct SignRequest {
    // message fields
    pub msg: ::std::string::String,
    // special fields
    pub unknown_fields: ::protobuf::UnknownFields,
    pub cached_size: ::protobuf::CachedSize,
}

impl<'a> ::std::default::Default for &'a SignRequest {
    fn default() -> &'a SignRequest {
        <SignRequest as ::protobuf::Message>::default_instance()
    }
}

impl SignRequest {
    pub fn new() -> SignRequest {
        ::std::default::Default::default()
    }

    // string msg = 1;


    pub fn get_msg(&self) -> &str {
        &self.msg
    }
    pub fn clear_msg(&mut self) {
        self.msg.clear();
    }

    // Param is passed by value, moved
    pub fn set_msg(&mut self, v: ::std::string::String) {
        self.msg = v;
    }

    // Mutable pointer to the field.
    // If field is not initialized, it is initialized with default value first.
    pub fn mut_msg(&mut self) -> &mut ::std::string::String {
        &mut self.msg
    }

    // Take field
    pub fn take_msg(&mut self) -> ::std::string::String {
        ::std::mem::replace(&mut self.msg, ::std::string::String::new())
    }
}

impl ::protobuf::Message for SignRequest {
    fn is_initialized(&self) -> bool {
        true
    }

    fn merge_from(&mut self, is: &mut ::protobuf::CodedInputStream<'_>) -> ::protobuf::ProtobufResult<()> {
        while !is.eof()? {
            let (field_number, wire_type) = is.read_tag_unpack()?;
            match field_number {
                1 => {
                    ::protobuf::rt::read_singular_proto3_string_into(wire_type, is, &mut self.msg)?;
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
        if !self.msg.is_empty() {
            my_size += ::protobuf::rt::string_size(1, &self.msg);
        }
        my_size += ::protobuf::rt::unknown_fields_size(self.get_unknown_fields());
        self.cached_size.set(my_size);
        my_size
    }

    fn write_to_with_cached_sizes(&self, os: &mut ::protobuf::CodedOutputStream<'_>) -> ::protobuf::ProtobufResult<()> {
        if !self.msg.is_empty() {
            os.write_string(1, &self.msg)?;
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

    fn new() -> SignRequest {
        SignRequest::new()
    }

    fn descriptor_static() -> &'static ::protobuf::reflect::MessageDescriptor {
        static descriptor: ::protobuf::rt::LazyV2<::protobuf::reflect::MessageDescriptor> = ::protobuf::rt::LazyV2::INIT;
        descriptor.get(|| {
            let mut fields = ::std::vec::Vec::new();
            fields.push(::protobuf::reflect::accessor::make_simple_field_accessor::<_, ::protobuf::types::ProtobufTypeString>(
                "msg",
                |m: &SignRequest| { &m.msg },
                |m: &mut SignRequest| { &mut m.msg },
            ));
            ::protobuf::reflect::MessageDescriptor::new_pb_name::<SignRequest>(
                "SignRequest",
                fields,
                file_descriptor_proto()
            )
        })
    }

    fn default_instance() -> &'static SignRequest {
        static instance: ::protobuf::rt::LazyV2<SignRequest> = ::protobuf::rt::LazyV2::INIT;
        instance.get(SignRequest::new)
    }
}

impl ::protobuf::Clear for SignRequest {
    fn clear(&mut self) {
        self.msg.clear();
        self.unknown_fields.clear();
    }
}

impl ::std::fmt::Debug for SignRequest {
    fn fmt(&self, f: &mut ::std::fmt::Formatter<'_>) -> ::std::fmt::Result {
        ::protobuf::text_format::fmt(self, f)
    }
}

impl ::protobuf::reflect::ProtobufValue for SignRequest {
    fn as_ref(&self) -> ::protobuf::reflect::ReflectValueRef {
        ::protobuf::reflect::ReflectValueRef::Message(self)
    }
}

#[derive(PartialEq,Clone,Default)]
pub struct SignResponse {
    // message fields
    pub signature: ::std::string::String,
    // special fields
    pub unknown_fields: ::protobuf::UnknownFields,
    pub cached_size: ::protobuf::CachedSize,
}

impl<'a> ::std::default::Default for &'a SignResponse {
    fn default() -> &'a SignResponse {
        <SignResponse as ::protobuf::Message>::default_instance()
    }
}

impl SignResponse {
    pub fn new() -> SignResponse {
        ::std::default::Default::default()
    }

    // string signature = 1;


    pub fn get_signature(&self) -> &str {
        &self.signature
    }
    pub fn clear_signature(&mut self) {
        self.signature.clear();
    }

    // Param is passed by value, moved
    pub fn set_signature(&mut self, v: ::std::string::String) {
        self.signature = v;
    }

    // Mutable pointer to the field.
    // If field is not initialized, it is initialized with default value first.
    pub fn mut_signature(&mut self) -> &mut ::std::string::String {
        &mut self.signature
    }

    // Take field
    pub fn take_signature(&mut self) -> ::std::string::String {
        ::std::mem::replace(&mut self.signature, ::std::string::String::new())
    }
}

impl ::protobuf::Message for SignResponse {
    fn is_initialized(&self) -> bool {
        true
    }

    fn merge_from(&mut self, is: &mut ::protobuf::CodedInputStream<'_>) -> ::protobuf::ProtobufResult<()> {
        while !is.eof()? {
            let (field_number, wire_type) = is.read_tag_unpack()?;
            match field_number {
                1 => {
                    ::protobuf::rt::read_singular_proto3_string_into(wire_type, is, &mut self.signature)?;
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
        if !self.signature.is_empty() {
            my_size += ::protobuf::rt::string_size(1, &self.signature);
        }
        my_size += ::protobuf::rt::unknown_fields_size(self.get_unknown_fields());
        self.cached_size.set(my_size);
        my_size
    }

    fn write_to_with_cached_sizes(&self, os: &mut ::protobuf::CodedOutputStream<'_>) -> ::protobuf::ProtobufResult<()> {
        if !self.signature.is_empty() {
            os.write_string(1, &self.signature)?;
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

    fn new() -> SignResponse {
        SignResponse::new()
    }

    fn descriptor_static() -> &'static ::protobuf::reflect::MessageDescriptor {
        static descriptor: ::protobuf::rt::LazyV2<::protobuf::reflect::MessageDescriptor> = ::protobuf::rt::LazyV2::INIT;
        descriptor.get(|| {
            let mut fields = ::std::vec::Vec::new();
            fields.push(::protobuf::reflect::accessor::make_simple_field_accessor::<_, ::protobuf::types::ProtobufTypeString>(
                "signature",
                |m: &SignResponse| { &m.signature },
                |m: &mut SignResponse| { &mut m.signature },
            ));
            ::protobuf::reflect::MessageDescriptor::new_pb_name::<SignResponse>(
                "SignResponse",
                fields,
                file_descriptor_proto()
            )
        })
    }

    fn default_instance() -> &'static SignResponse {
        static instance: ::protobuf::rt::LazyV2<SignResponse> = ::protobuf::rt::LazyV2::INIT;
        instance.get(SignResponse::new)
    }
}

impl ::protobuf::Clear for SignResponse {
    fn clear(&mut self) {
        self.signature.clear();
        self.unknown_fields.clear();
    }
}

impl ::std::fmt::Debug for SignResponse {
    fn fmt(&self, f: &mut ::std::fmt::Formatter<'_>) -> ::std::fmt::Result {
        ::protobuf::text_format::fmt(self, f)
    }
}

impl ::protobuf::reflect::ProtobufValue for SignResponse {
    fn as_ref(&self) -> ::protobuf::reflect::ReflectValueRef {
        ::protobuf::reflect::ReflectValueRef::Message(self)
    }
}

#[derive(PartialEq,Clone,Default)]
pub struct VerifyRequest {
    // message fields
    pub msg: ::std::string::String,
    pub signature: ::std::string::String,
    // special fields
    pub unknown_fields: ::protobuf::UnknownFields,
    pub cached_size: ::protobuf::CachedSize,
}

impl<'a> ::std::default::Default for &'a VerifyRequest {
    fn default() -> &'a VerifyRequest {
        <VerifyRequest as ::protobuf::Message>::default_instance()
    }
}

impl VerifyRequest {
    pub fn new() -> VerifyRequest {
        ::std::default::Default::default()
    }

    // string msg = 1;


    pub fn get_msg(&self) -> &str {
        &self.msg
    }
    pub fn clear_msg(&mut self) {
        self.msg.clear();
    }

    // Param is passed by value, moved
    pub fn set_msg(&mut self, v: ::std::string::String) {
        self.msg = v;
    }

    // Mutable pointer to the field.
    // If field is not initialized, it is initialized with default value first.
    pub fn mut_msg(&mut self) -> &mut ::std::string::String {
        &mut self.msg
    }

    // Take field
    pub fn take_msg(&mut self) -> ::std::string::String {
        ::std::mem::replace(&mut self.msg, ::std::string::String::new())
    }

    // string signature = 2;


    pub fn get_signature(&self) -> &str {
        &self.signature
    }
    pub fn clear_signature(&mut self) {
        self.signature.clear();
    }

    // Param is passed by value, moved
    pub fn set_signature(&mut self, v: ::std::string::String) {
        self.signature = v;
    }

    // Mutable pointer to the field.
    // If field is not initialized, it is initialized with default value first.
    pub fn mut_signature(&mut self) -> &mut ::std::string::String {
        &mut self.signature
    }

    // Take field
    pub fn take_signature(&mut self) -> ::std::string::String {
        ::std::mem::replace(&mut self.signature, ::std::string::String::new())
    }
}

impl ::protobuf::Message for VerifyRequest {
    fn is_initialized(&self) -> bool {
        true
    }

    fn merge_from(&mut self, is: &mut ::protobuf::CodedInputStream<'_>) -> ::protobuf::ProtobufResult<()> {
        while !is.eof()? {
            let (field_number, wire_type) = is.read_tag_unpack()?;
            match field_number {
                1 => {
                    ::protobuf::rt::read_singular_proto3_string_into(wire_type, is, &mut self.msg)?;
                },
                2 => {
                    ::protobuf::rt::read_singular_proto3_string_into(wire_type, is, &mut self.signature)?;
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
        if !self.msg.is_empty() {
            my_size += ::protobuf::rt::string_size(1, &self.msg);
        }
        if !self.signature.is_empty() {
            my_size += ::protobuf::rt::string_size(2, &self.signature);
        }
        my_size += ::protobuf::rt::unknown_fields_size(self.get_unknown_fields());
        self.cached_size.set(my_size);
        my_size
    }

    fn write_to_with_cached_sizes(&self, os: &mut ::protobuf::CodedOutputStream<'_>) -> ::protobuf::ProtobufResult<()> {
        if !self.msg.is_empty() {
            os.write_string(1, &self.msg)?;
        }
        if !self.signature.is_empty() {
            os.write_string(2, &self.signature)?;
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

    fn new() -> VerifyRequest {
        VerifyRequest::new()
    }

    fn descriptor_static() -> &'static ::protobuf::reflect::MessageDescriptor {
        static descriptor: ::protobuf::rt::LazyV2<::protobuf::reflect::MessageDescriptor> = ::protobuf::rt::LazyV2::INIT;
        descriptor.get(|| {
            let mut fields = ::std::vec::Vec::new();
            fields.push(::protobuf::reflect::accessor::make_simple_field_accessor::<_, ::protobuf::types::ProtobufTypeString>(
                "msg",
                |m: &VerifyRequest| { &m.msg },
                |m: &mut VerifyRequest| { &mut m.msg },
            ));
            fields.push(::protobuf::reflect::accessor::make_simple_field_accessor::<_, ::protobuf::types::ProtobufTypeString>(
                "signature",
                |m: &VerifyRequest| { &m.signature },
                |m: &mut VerifyRequest| { &mut m.signature },
            ));
            ::protobuf::reflect::MessageDescriptor::new_pb_name::<VerifyRequest>(
                "VerifyRequest",
                fields,
                file_descriptor_proto()
            )
        })
    }

    fn default_instance() -> &'static VerifyRequest {
        static instance: ::protobuf::rt::LazyV2<VerifyRequest> = ::protobuf::rt::LazyV2::INIT;
        instance.get(VerifyRequest::new)
    }
}

impl ::protobuf::Clear for VerifyRequest {
    fn clear(&mut self) {
        self.msg.clear();
        self.signature.clear();
        self.unknown_fields.clear();
    }
}

impl ::std::fmt::Debug for VerifyRequest {
    fn fmt(&self, f: &mut ::std::fmt::Formatter<'_>) -> ::std::fmt::Result {
        ::protobuf::text_format::fmt(self, f)
    }
}

impl ::protobuf::reflect::ProtobufValue for VerifyRequest {
    fn as_ref(&self) -> ::protobuf::reflect::ReflectValueRef {
        ::protobuf::reflect::ReflectValueRef::Message(self)
    }
}

#[derive(PartialEq,Clone,Default)]
pub struct VerifyResponse {
    // message fields
    pub result: bool,
    // special fields
    pub unknown_fields: ::protobuf::UnknownFields,
    pub cached_size: ::protobuf::CachedSize,
}

impl<'a> ::std::default::Default for &'a VerifyResponse {
    fn default() -> &'a VerifyResponse {
        <VerifyResponse as ::protobuf::Message>::default_instance()
    }
}

impl VerifyResponse {
    pub fn new() -> VerifyResponse {
        ::std::default::Default::default()
    }

    // bool result = 1;


    pub fn get_result(&self) -> bool {
        self.result
    }
    pub fn clear_result(&mut self) {
        self.result = false;
    }

    // Param is passed by value, moved
    pub fn set_result(&mut self, v: bool) {
        self.result = v;
    }
}

impl ::protobuf::Message for VerifyResponse {
    fn is_initialized(&self) -> bool {
        true
    }

    fn merge_from(&mut self, is: &mut ::protobuf::CodedInputStream<'_>) -> ::protobuf::ProtobufResult<()> {
        while !is.eof()? {
            let (field_number, wire_type) = is.read_tag_unpack()?;
            match field_number {
                1 => {
                    if wire_type != ::protobuf::wire_format::WireTypeVarint {
                        return ::std::result::Result::Err(::protobuf::rt::unexpected_wire_type(wire_type));
                    }
                    let tmp = is.read_bool()?;
                    self.result = tmp;
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
        if self.result != false {
            my_size += 2;
        }
        my_size += ::protobuf::rt::unknown_fields_size(self.get_unknown_fields());
        self.cached_size.set(my_size);
        my_size
    }

    fn write_to_with_cached_sizes(&self, os: &mut ::protobuf::CodedOutputStream<'_>) -> ::protobuf::ProtobufResult<()> {
        if self.result != false {
            os.write_bool(1, self.result)?;
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

    fn new() -> VerifyResponse {
        VerifyResponse::new()
    }

    fn descriptor_static() -> &'static ::protobuf::reflect::MessageDescriptor {
        static descriptor: ::protobuf::rt::LazyV2<::protobuf::reflect::MessageDescriptor> = ::protobuf::rt::LazyV2::INIT;
        descriptor.get(|| {
            let mut fields = ::std::vec::Vec::new();
            fields.push(::protobuf::reflect::accessor::make_simple_field_accessor::<_, ::protobuf::types::ProtobufTypeBool>(
                "result",
                |m: &VerifyResponse| { &m.result },
                |m: &mut VerifyResponse| { &mut m.result },
            ));
            ::protobuf::reflect::MessageDescriptor::new_pb_name::<VerifyResponse>(
                "VerifyResponse",
                fields,
                file_descriptor_proto()
            )
        })
    }

    fn default_instance() -> &'static VerifyResponse {
        static instance: ::protobuf::rt::LazyV2<VerifyResponse> = ::protobuf::rt::LazyV2::INIT;
        instance.get(VerifyResponse::new)
    }
}

impl ::protobuf::Clear for VerifyResponse {
    fn clear(&mut self) {
        self.result = false;
        self.unknown_fields.clear();
    }
}

impl ::std::fmt::Debug for VerifyResponse {
    fn fmt(&self, f: &mut ::std::fmt::Formatter<'_>) -> ::std::fmt::Result {
        ::protobuf::text_format::fmt(self, f)
    }
}

impl ::protobuf::reflect::ProtobufValue for VerifyResponse {
    fn as_ref(&self) -> ::protobuf::reflect::ReflectValueRef {
        ::protobuf::reflect::ReflectValueRef::Message(self)
    }
}

static file_descriptor_proto_data: &'static [u8] = b"\
    \n#proto/imp/api/signing/signing.proto\x12\x07signing\x1a\x1cgoogle/api/\
    annotations.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"\
    \x1f\n\x0bSignRequest\x12\x10\n\x03msg\x18\x01\x20\x01(\tR\x03msg\",\n\
    \x0cSignResponse\x12\x1c\n\tsignature\x18\x01\x20\x01(\tR\tsignature\"?\
    \n\rVerifyRequest\x12\x10\n\x03msg\x18\x01\x20\x01(\tR\x03msg\x12\x1c\n\
    \tsignature\x18\x02\x20\x01(\tR\tsignature\"(\n\x0eVerifyResponse\x12\
    \x16\n\x06result\x18\x01\x20\x01(\x08R\x06result2\xb5\x01\n\x07Signing\
    \x12O\n\x0bSignMessage\x12\x14.signing.SignRequest\x1a\x15.signing.SignR\
    esponse\"\x13\x82\xd3\xe4\x93\x02\r\"\x08/v1/sign:\x01*\x12Y\n\x0fVerify\
    Signature\x12\x16.signing.VerifyRequest\x1a\x17.signing.VerifyResponse\"\
    \x15\x82\xd3\xe4\x93\x02\x0f\"\n/v1/verify:\x01*B\xc7\x01Z#github.com/im\
    perviousai/freeimp/gen\x92A\x9e\x01\x12?\n\x10Signing\x20Services\"&\n\r\
    Impervious\x20AI\x12\x15https://impervious.ai2\x031.0*\x03\x01\x02\x042\
    \x10application/json:\x10application/jsonr2\n\x14Documentation\x20on\x20\
    IMP\x12\x1ahttps://docs.impervious.aiJ\xb8\n\n\x06\x12\x04\x01\0R\x01\nS\
    \n\x01\x0c\x12\x03\x01\0\x12\x1aI/\x20Allows\x20an\x20Imp\x20node\x20to\
    \x20sign\x20and\x20verify\x20messaging\x20with\x20the\x20connected\x20LN\
    D\n\n\x08\n\x01\x02\x12\x03\x03\0\x10\n\x08\n\x01\x08\x12\x03\x05\0:\n\t\
    \n\x02\x08\x0b\x12\x03\x05\0:\n\t\n\x02\x03\0\x12\x03\x07\0&\n\t\n\x02\
    \x03\x01\x12\x03\x08\08\n\t\n\x01\x08\x12\x04\n\0\x1c\x02\n\x0b\n\x03\
    \x08\x92\x08\x12\x04\n\0\x1c\x02\nq\n\x02\x06\0\x12\x04!\05\x01\x1ae*\n\
    \x20Signing\x20service\x20allows\x20an\x20Imp\x20node\x20to\x20sign\x20a\
    nd\x20verify\x20messages\x20with\x20the\x20connected\x20lightning\x20nod\
    e.\n\n\n\n\x03\x06\0\x01\x12\x03!\x08\x0f\nK\n\x04\x06\0\x02\0\x12\x04%\
    \x08*\t\x1a=*\n\x20SignMessage\x20signs\x20a\x20message\x20with\x20your\
    \x20node's\x20private\x20key.\n\n\x0c\n\x05\x06\0\x02\0\x01\x12\x03%\x0c\
    \x17\n\x0c\n\x05\x06\0\x02\0\x02\x12\x03%\x18#\n\x0c\n\x05\x06\0\x02\0\
    \x03\x12\x03%.:\n\r\n\x05\x06\0\x02\0\x04\x12\x04&\x10)\x12\n\x11\n\t\
    \x06\0\x02\0\x04\xb0\xca\xbc\"\x12\x04&\x10)\x12\nP\n\x04\x06\0\x02\x01\
    \x12\x04/\x084\t\x1aB*\n\x20Verifymessage\x20verifies\x20a\x20message\
    \x20was\x20signed\x20from\x20another\x20node.\n\n\x0c\n\x05\x06\0\x02\
    \x01\x01\x12\x03/\x0c\x1b\n\x0c\n\x05\x06\0\x02\x01\x02\x12\x03/\x1c)\n\
    \x0c\n\x05\x06\0\x02\x01\x03\x12\x03/4B\n\r\n\x05\x06\0\x02\x01\x04\x12\
    \x040\x103\x12\n\x11\n\t\x06\0\x02\x01\x04\xb0\xca\xbc\"\x12\x040\x103\
    \x12\n6\n\x02\x04\0\x12\x04:\0<\x01\x1a**\n\x20Represents\x20a\x20reques\
    t\x20to\x20sign\x20a\x20message\n\n\n\n\x03\x04\0\x01\x12\x03:\x08\x13\n\
    #\n\x04\x04\0\x02\0\x12\x03;\x08\x17\"\x16\x20message\x20to\x20be\x20sig\
    ned\n\n\x0c\n\x05\x04\0\x02\0\x05\x12\x03;\x08\x0e\n\x0c\n\x05\x04\0\x02\
    \0\x01\x12\x03;\x0f\x12\n\x0c\n\x05\x04\0\x02\0\x03\x12\x03;\x15\x16\n>\
    \n\x02\x04\x01\x12\x04A\0C\x01\x1a2*\n\x20Represents\x20a\x20response\
    \x20from\x20a\x20signature\x20request\n\n\n\n\x03\x04\x01\x01\x12\x03A\
    \x08\x14\n*\n\x04\x04\x01\x02\0\x12\x03B\x08\x1d\"\x1d\x20signature\x20o\
    f\x20signed\x20message\n\n\x0c\n\x05\x04\x01\x02\0\x05\x12\x03B\x08\x0e\
    \n\x0c\n\x05\x04\x01\x02\0\x01\x12\x03B\x0f\x18\n\x0c\n\x05\x04\x01\x02\
    \0\x03\x12\x03B\x1b\x1c\nF\n\x02\x04\x02\x12\x04H\0K\x01\x1a:*\n\x20Repr\
    esents\x20a\x20request\x20to\x20verify\x20a\x20signature\x20and\x20messa\
    ge\n\n\n\n\x03\x04\x02\x01\x12\x03H\x08\x15\n%\n\x04\x04\x02\x02\0\x12\
    \x03I\x08\x17\"\x18\x20message\x20to\x20be\x20verified\n\n\x0c\n\x05\x04\
    \x02\x02\0\x05\x12\x03I\x08\x0e\n\x0c\n\x05\x04\x02\x02\0\x01\x12\x03I\
    \x0f\x12\n\x0c\n\x05\x04\x02\x02\0\x03\x12\x03I\x15\x16\n#\n\x04\x04\x02\
    \x02\x01\x12\x03J\x08\x1d\"\x16\x20signature\x20of\x20message\n\n\x0c\n\
    \x05\x04\x02\x02\x01\x05\x12\x03J\x08\x0e\n\x0c\n\x05\x04\x02\x02\x01\
    \x01\x12\x03J\x0f\x18\n\x0c\n\x05\x04\x02\x02\x01\x03\x12\x03J\x1b\x1c\n\
    F\n\x02\x04\x03\x12\x04P\0R\x01\x1a:*\n\x20Represents\x20a\x20response\
    \x20back\x20from\x20a\x20verification\x20request\n\n\n\n\x03\x04\x03\x01\
    \x12\x03P\x08\x16\n/\n\x04\x04\x03\x02\0\x12\x03Q\x08\x18\"\"\x20result\
    \x20of\x20signature\x20verification\n\n\x0c\n\x05\x04\x03\x02\0\x05\x12\
    \x03Q\x08\x0c\n\x0c\n\x05\x04\x03\x02\0\x01\x12\x03Q\r\x13\n\x0c\n\x05\
    \x04\x03\x02\0\x03\x12\x03Q\x16\x17b\x06proto3\
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
