# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: maze.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='maze.proto',
  package='maze',
  syntax='proto3',
  serialized_options=b'Z\006mazepb',
  serialized_pb=b'\n\nmaze.proto\x12\x04maze\"\"\n\nGenRequest\x12\t\n\x01w\x18\x01 \x01(\x11\x12\t\n\x01h\x18\x02 \x01(\x11\"\x1c\n\x0bGenResponse\x12\r\n\x05\x66ield\x18\x01 \x01(\t22\n\x04Maze\x12*\n\x03Gen\x12\x10.maze.GenRequest\x1a\x11.maze.GenResponseB\x08Z\x06mazepbb\x06proto3'
)




_GENREQUEST = _descriptor.Descriptor(
  name='GenRequest',
  full_name='maze.GenRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='w', full_name='maze.GenRequest.w', index=0,
      number=1, type=17, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='h', full_name='maze.GenRequest.h', index=1,
      number=2, type=17, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=20,
  serialized_end=54,
)


_GENRESPONSE = _descriptor.Descriptor(
  name='GenResponse',
  full_name='maze.GenResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='field', full_name='maze.GenResponse.field', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=56,
  serialized_end=84,
)

DESCRIPTOR.message_types_by_name['GenRequest'] = _GENREQUEST
DESCRIPTOR.message_types_by_name['GenResponse'] = _GENRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

GenRequest = _reflection.GeneratedProtocolMessageType('GenRequest', (_message.Message,), {
  'DESCRIPTOR' : _GENREQUEST,
  '__module__' : 'maze_pb2'
  # @@protoc_insertion_point(class_scope:maze.GenRequest)
  })
_sym_db.RegisterMessage(GenRequest)

GenResponse = _reflection.GeneratedProtocolMessageType('GenResponse', (_message.Message,), {
  'DESCRIPTOR' : _GENRESPONSE,
  '__module__' : 'maze_pb2'
  # @@protoc_insertion_point(class_scope:maze.GenResponse)
  })
_sym_db.RegisterMessage(GenResponse)


DESCRIPTOR._options = None

_MAZE = _descriptor.ServiceDescriptor(
  name='Maze',
  full_name='maze.Maze',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=86,
  serialized_end=136,
  methods=[
  _descriptor.MethodDescriptor(
    name='Gen',
    full_name='maze.Maze.Gen',
    index=0,
    containing_service=None,
    input_type=_GENREQUEST,
    output_type=_GENRESPONSE,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_MAZE)

DESCRIPTOR.services_by_name['Maze'] = _MAZE

# @@protoc_insertion_point(module_scope)
