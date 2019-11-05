# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: controlservice.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from . import datamodel_pb2 as datamodel__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='controlservice.proto',
  package='service',
  syntax='proto3',
  serialized_options=_b('\n\026org.qmstr.grpc.service'),
  serialized_pb=_b('\n\x14\x63ontrolservice.proto\x12\x07service\x1a\x0f\x64\x61tamodel.proto\"\x19\n\nLogMessage\x12\x0b\n\x03msg\x18\x01 \x01(\x0c\"\x1e\n\x0bLogResponse\x12\x0f\n\x07success\x18\x01 \x01(\x08\"\x1b\n\x0bQuitMessage\x12\x0c\n\x04kill\x18\x01 \x01(\x08\"\x1f\n\x0cQuitResponse\x12\x0f\n\x07success\x18\x01 \x01(\x08\"3\n\x12SwitchPhaseMessage\x12\x1d\n\x05phase\x18\x01 \x01(\x0e\x32\x0e.service.Phase\"5\n\x13SwitchPhaseResponse\x12\x0f\n\x07success\x18\x01 \x01(\x08\x12\r\n\x05\x65rror\x18\x02 \x01(\t\"M\n\x12GetFileNodeMessage\x12#\n\x08\x66ileNode\x18\x01 \x01(\x0b\x32\x11.service.FileNode\x12\x12\n\nuniqueNode\x18\x02 \x01(\x08\".\n\rStatusMessage\x12\r\n\x05phase\x18\x01 \x01(\x08\x12\x0e\n\x06switch\x18\x02 \x01(\x08\"z\n\x0eStatusResponse\x12\r\n\x05phase\x18\x01 \x01(\t\x12\x1f\n\x07phaseID\x18\x02 \x01(\x0e\x32\x0e.service.Phase\x12\x11\n\tswitching\x18\x03 \x01(\x08\x12\r\n\x05\x65rror\x18\x04 \x01(\t\x12\x16\n\x0ependingInserts\x18\x05 \x01(\x04\"2\n\x0c\x45ventMessage\x12\"\n\x05\x63lass\x18\x01 \x01(\x0e\x32\x13.service.EventClass\"\x1d\n\rExportRequest\x12\x0c\n\x04wait\x18\x01 \x01(\x08\"!\n\x0e\x45xportResponse\x12\x0f\n\x07success\x18\x01 \x01(\x08\x32\x99\x05\n\x0e\x43ontrolService\x12\x32\n\x03Log\x12\x13.service.LogMessage\x1a\x14.service.LogResponse\"\x00\x12\x35\n\x04Quit\x12\x14.service.QuitMessage\x1a\x15.service.QuitResponse\"\x00\x12J\n\x0bSwitchPhase\x12\x1b.service.SwitchPhaseMessage\x1a\x1c.service.SwitchPhaseResponse\"\x00\x12@\n\x0eGetPackageNode\x12\x14.service.PackageNode\x1a\x14.service.PackageNode\"\x00\x30\x01\x12@\n\x11GetPackageTargets\x12\x14.service.PackageNode\x1a\x11.service.FileNode\"\x00\x30\x01\x12\x41\n\x0bGetFileNode\x12\x1b.service.GetFileNodeMessage\x1a\x11.service.FileNode\"\x00\x30\x01\x12I\n\x11GetDiagnosticNode\x12\x17.service.DiagnosticNode\x1a\x17.service.DiagnosticNode\"\x00\x30\x01\x12;\n\x06Status\x12\x16.service.StatusMessage\x1a\x17.service.StatusResponse\"\x00\x12<\n\x0fSubscribeEvents\x12\x15.service.EventMessage\x1a\x0e.service.Event\"\x00\x30\x01\x12\x43\n\x0e\x45xportSnapshot\x12\x16.service.ExportRequest\x1a\x17.service.ExportResponse\"\x00\x42\x18\n\x16org.qmstr.grpc.serviceX\x00\x62\x06proto3')
  ,
  dependencies=[datamodel__pb2.DESCRIPTOR,])




_LOGMESSAGE = _descriptor.Descriptor(
  name='LogMessage',
  full_name='service.LogMessage',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='msg', full_name='service.LogMessage.msg', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
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
  serialized_start=50,
  serialized_end=75,
)


_LOGRESPONSE = _descriptor.Descriptor(
  name='LogResponse',
  full_name='service.LogResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='success', full_name='service.LogResponse.success', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
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
  serialized_start=77,
  serialized_end=107,
)


_QUITMESSAGE = _descriptor.Descriptor(
  name='QuitMessage',
  full_name='service.QuitMessage',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='kill', full_name='service.QuitMessage.kill', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
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
  serialized_start=109,
  serialized_end=136,
)


_QUITRESPONSE = _descriptor.Descriptor(
  name='QuitResponse',
  full_name='service.QuitResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='success', full_name='service.QuitResponse.success', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
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
  serialized_start=138,
  serialized_end=169,
)


_SWITCHPHASEMESSAGE = _descriptor.Descriptor(
  name='SwitchPhaseMessage',
  full_name='service.SwitchPhaseMessage',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='phase', full_name='service.SwitchPhaseMessage.phase', index=0,
      number=1, type=14, cpp_type=8, label=1,
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
  serialized_start=171,
  serialized_end=222,
)


_SWITCHPHASERESPONSE = _descriptor.Descriptor(
  name='SwitchPhaseResponse',
  full_name='service.SwitchPhaseResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='success', full_name='service.SwitchPhaseResponse.success', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='error', full_name='service.SwitchPhaseResponse.error', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
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
  serialized_start=224,
  serialized_end=277,
)


_GETFILENODEMESSAGE = _descriptor.Descriptor(
  name='GetFileNodeMessage',
  full_name='service.GetFileNodeMessage',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='fileNode', full_name='service.GetFileNodeMessage.fileNode', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='uniqueNode', full_name='service.GetFileNodeMessage.uniqueNode', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
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
  serialized_start=279,
  serialized_end=356,
)


_STATUSMESSAGE = _descriptor.Descriptor(
  name='StatusMessage',
  full_name='service.StatusMessage',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='phase', full_name='service.StatusMessage.phase', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='switch', full_name='service.StatusMessage.switch', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
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
  serialized_start=358,
  serialized_end=404,
)


_STATUSRESPONSE = _descriptor.Descriptor(
  name='StatusResponse',
  full_name='service.StatusResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='phase', full_name='service.StatusResponse.phase', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='phaseID', full_name='service.StatusResponse.phaseID', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='switching', full_name='service.StatusResponse.switching', index=2,
      number=3, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='error', full_name='service.StatusResponse.error', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='pendingInserts', full_name='service.StatusResponse.pendingInserts', index=4,
      number=5, type=4, cpp_type=4, label=1,
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
  serialized_start=406,
  serialized_end=528,
)


_EVENTMESSAGE = _descriptor.Descriptor(
  name='EventMessage',
  full_name='service.EventMessage',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='class', full_name='service.EventMessage.class', index=0,
      number=1, type=14, cpp_type=8, label=1,
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
  serialized_start=530,
  serialized_end=580,
)


_EXPORTREQUEST = _descriptor.Descriptor(
  name='ExportRequest',
  full_name='service.ExportRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='wait', full_name='service.ExportRequest.wait', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
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
  serialized_start=582,
  serialized_end=611,
)


_EXPORTRESPONSE = _descriptor.Descriptor(
  name='ExportResponse',
  full_name='service.ExportResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='success', full_name='service.ExportResponse.success', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
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
  serialized_start=613,
  serialized_end=646,
)

_SWITCHPHASEMESSAGE.fields_by_name['phase'].enum_type = datamodel__pb2._PHASE
_GETFILENODEMESSAGE.fields_by_name['fileNode'].message_type = datamodel__pb2._FILENODE
_STATUSRESPONSE.fields_by_name['phaseID'].enum_type = datamodel__pb2._PHASE
_EVENTMESSAGE.fields_by_name['class'].enum_type = datamodel__pb2._EVENTCLASS
DESCRIPTOR.message_types_by_name['LogMessage'] = _LOGMESSAGE
DESCRIPTOR.message_types_by_name['LogResponse'] = _LOGRESPONSE
DESCRIPTOR.message_types_by_name['QuitMessage'] = _QUITMESSAGE
DESCRIPTOR.message_types_by_name['QuitResponse'] = _QUITRESPONSE
DESCRIPTOR.message_types_by_name['SwitchPhaseMessage'] = _SWITCHPHASEMESSAGE
DESCRIPTOR.message_types_by_name['SwitchPhaseResponse'] = _SWITCHPHASERESPONSE
DESCRIPTOR.message_types_by_name['GetFileNodeMessage'] = _GETFILENODEMESSAGE
DESCRIPTOR.message_types_by_name['StatusMessage'] = _STATUSMESSAGE
DESCRIPTOR.message_types_by_name['StatusResponse'] = _STATUSRESPONSE
DESCRIPTOR.message_types_by_name['EventMessage'] = _EVENTMESSAGE
DESCRIPTOR.message_types_by_name['ExportRequest'] = _EXPORTREQUEST
DESCRIPTOR.message_types_by_name['ExportResponse'] = _EXPORTRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

LogMessage = _reflection.GeneratedProtocolMessageType('LogMessage', (_message.Message,), dict(
  DESCRIPTOR = _LOGMESSAGE,
  __module__ = 'controlservice_pb2'
  # @@protoc_insertion_point(class_scope:service.LogMessage)
  ))
_sym_db.RegisterMessage(LogMessage)

LogResponse = _reflection.GeneratedProtocolMessageType('LogResponse', (_message.Message,), dict(
  DESCRIPTOR = _LOGRESPONSE,
  __module__ = 'controlservice_pb2'
  # @@protoc_insertion_point(class_scope:service.LogResponse)
  ))
_sym_db.RegisterMessage(LogResponse)

QuitMessage = _reflection.GeneratedProtocolMessageType('QuitMessage', (_message.Message,), dict(
  DESCRIPTOR = _QUITMESSAGE,
  __module__ = 'controlservice_pb2'
  # @@protoc_insertion_point(class_scope:service.QuitMessage)
  ))
_sym_db.RegisterMessage(QuitMessage)

QuitResponse = _reflection.GeneratedProtocolMessageType('QuitResponse', (_message.Message,), dict(
  DESCRIPTOR = _QUITRESPONSE,
  __module__ = 'controlservice_pb2'
  # @@protoc_insertion_point(class_scope:service.QuitResponse)
  ))
_sym_db.RegisterMessage(QuitResponse)

SwitchPhaseMessage = _reflection.GeneratedProtocolMessageType('SwitchPhaseMessage', (_message.Message,), dict(
  DESCRIPTOR = _SWITCHPHASEMESSAGE,
  __module__ = 'controlservice_pb2'
  # @@protoc_insertion_point(class_scope:service.SwitchPhaseMessage)
  ))
_sym_db.RegisterMessage(SwitchPhaseMessage)

SwitchPhaseResponse = _reflection.GeneratedProtocolMessageType('SwitchPhaseResponse', (_message.Message,), dict(
  DESCRIPTOR = _SWITCHPHASERESPONSE,
  __module__ = 'controlservice_pb2'
  # @@protoc_insertion_point(class_scope:service.SwitchPhaseResponse)
  ))
_sym_db.RegisterMessage(SwitchPhaseResponse)

GetFileNodeMessage = _reflection.GeneratedProtocolMessageType('GetFileNodeMessage', (_message.Message,), dict(
  DESCRIPTOR = _GETFILENODEMESSAGE,
  __module__ = 'controlservice_pb2'
  # @@protoc_insertion_point(class_scope:service.GetFileNodeMessage)
  ))
_sym_db.RegisterMessage(GetFileNodeMessage)

StatusMessage = _reflection.GeneratedProtocolMessageType('StatusMessage', (_message.Message,), dict(
  DESCRIPTOR = _STATUSMESSAGE,
  __module__ = 'controlservice_pb2'
  # @@protoc_insertion_point(class_scope:service.StatusMessage)
  ))
_sym_db.RegisterMessage(StatusMessage)

StatusResponse = _reflection.GeneratedProtocolMessageType('StatusResponse', (_message.Message,), dict(
  DESCRIPTOR = _STATUSRESPONSE,
  __module__ = 'controlservice_pb2'
  # @@protoc_insertion_point(class_scope:service.StatusResponse)
  ))
_sym_db.RegisterMessage(StatusResponse)

EventMessage = _reflection.GeneratedProtocolMessageType('EventMessage', (_message.Message,), dict(
  DESCRIPTOR = _EVENTMESSAGE,
  __module__ = 'controlservice_pb2'
  # @@protoc_insertion_point(class_scope:service.EventMessage)
  ))
_sym_db.RegisterMessage(EventMessage)

ExportRequest = _reflection.GeneratedProtocolMessageType('ExportRequest', (_message.Message,), dict(
  DESCRIPTOR = _EXPORTREQUEST,
  __module__ = 'controlservice_pb2'
  # @@protoc_insertion_point(class_scope:service.ExportRequest)
  ))
_sym_db.RegisterMessage(ExportRequest)

ExportResponse = _reflection.GeneratedProtocolMessageType('ExportResponse', (_message.Message,), dict(
  DESCRIPTOR = _EXPORTRESPONSE,
  __module__ = 'controlservice_pb2'
  # @@protoc_insertion_point(class_scope:service.ExportResponse)
  ))
_sym_db.RegisterMessage(ExportResponse)


DESCRIPTOR._options = None

_CONTROLSERVICE = _descriptor.ServiceDescriptor(
  name='ControlService',
  full_name='service.ControlService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=649,
  serialized_end=1314,
  methods=[
  _descriptor.MethodDescriptor(
    name='Log',
    full_name='service.ControlService.Log',
    index=0,
    containing_service=None,
    input_type=_LOGMESSAGE,
    output_type=_LOGRESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='Quit',
    full_name='service.ControlService.Quit',
    index=1,
    containing_service=None,
    input_type=_QUITMESSAGE,
    output_type=_QUITRESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='SwitchPhase',
    full_name='service.ControlService.SwitchPhase',
    index=2,
    containing_service=None,
    input_type=_SWITCHPHASEMESSAGE,
    output_type=_SWITCHPHASERESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='GetPackageNode',
    full_name='service.ControlService.GetPackageNode',
    index=3,
    containing_service=None,
    input_type=datamodel__pb2._PACKAGENODE,
    output_type=datamodel__pb2._PACKAGENODE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='GetPackageTargets',
    full_name='service.ControlService.GetPackageTargets',
    index=4,
    containing_service=None,
    input_type=datamodel__pb2._PACKAGENODE,
    output_type=datamodel__pb2._FILENODE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='GetFileNode',
    full_name='service.ControlService.GetFileNode',
    index=5,
    containing_service=None,
    input_type=_GETFILENODEMESSAGE,
    output_type=datamodel__pb2._FILENODE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='GetDiagnosticNode',
    full_name='service.ControlService.GetDiagnosticNode',
    index=6,
    containing_service=None,
    input_type=datamodel__pb2._DIAGNOSTICNODE,
    output_type=datamodel__pb2._DIAGNOSTICNODE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='Status',
    full_name='service.ControlService.Status',
    index=7,
    containing_service=None,
    input_type=_STATUSMESSAGE,
    output_type=_STATUSRESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='SubscribeEvents',
    full_name='service.ControlService.SubscribeEvents',
    index=8,
    containing_service=None,
    input_type=_EVENTMESSAGE,
    output_type=datamodel__pb2._EVENT,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='ExportSnapshot',
    full_name='service.ControlService.ExportSnapshot',
    index=9,
    containing_service=None,
    input_type=_EXPORTREQUEST,
    output_type=_EXPORTRESPONSE,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_CONTROLSERVICE)

DESCRIPTOR.services_by_name['ControlService'] = _CONTROLSERVICE

# @@protoc_insertion_point(module_scope)
