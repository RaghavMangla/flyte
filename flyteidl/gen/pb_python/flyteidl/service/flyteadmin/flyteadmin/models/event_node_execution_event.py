# coding: utf-8

"""
    flyteidl/service/admin.proto

    No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)  # noqa: E501

    OpenAPI spec version: version not set
    
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


import pprint
import re  # noqa: F401

import six

from flyteadmin.models.core_execution_error import CoreExecutionError  # noqa: F401,E501
from flyteadmin.models.core_literal_map import CoreLiteralMap  # noqa: F401,E501
from flyteadmin.models.core_node_execution_identifier import CoreNodeExecutionIdentifier  # noqa: F401,E501
from flyteadmin.models.core_node_execution_phase import CoreNodeExecutionPhase  # noqa: F401,E501
from flyteadmin.models.event_parent_node_execution_metadata import EventParentNodeExecutionMetadata  # noqa: F401,E501
from flyteadmin.models.event_parent_task_execution_metadata import EventParentTaskExecutionMetadata  # noqa: F401,E501
from flyteadmin.models.flyteidlevent_task_node_metadata import FlyteidleventTaskNodeMetadata  # noqa: F401,E501
from flyteadmin.models.flyteidlevent_workflow_node_metadata import FlyteidleventWorkflowNodeMetadata  # noqa: F401,E501


class EventNodeExecutionEvent(object):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    """
    Attributes:
      swagger_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    swagger_types = {
        'id': 'CoreNodeExecutionIdentifier',
        'producer_id': 'str',
        'phase': 'CoreNodeExecutionPhase',
        'occurred_at': 'datetime',
        'input_uri': 'str',
        'input_data': 'CoreLiteralMap',
        'output_uri': 'str',
        'error': 'CoreExecutionError',
        'output_data': 'CoreLiteralMap',
        'workflow_node_metadata': 'FlyteidleventWorkflowNodeMetadata',
        'task_node_metadata': 'FlyteidleventTaskNodeMetadata',
        'parent_task_metadata': 'EventParentTaskExecutionMetadata',
        'parent_node_metadata': 'EventParentNodeExecutionMetadata',
        'retry_group': 'str',
        'spec_node_id': 'str',
        'node_name': 'str',
        'event_version': 'int',
        'is_parent': 'bool',
        'is_dynamic': 'bool',
        'deck_uri': 'str',
        'reported_at': 'datetime'
    }

    attribute_map = {
        'id': 'id',
        'producer_id': 'producer_id',
        'phase': 'phase',
        'occurred_at': 'occurred_at',
        'input_uri': 'input_uri',
        'input_data': 'input_data',
        'output_uri': 'output_uri',
        'error': 'error',
        'output_data': 'output_data',
        'workflow_node_metadata': 'workflow_node_metadata',
        'task_node_metadata': 'task_node_metadata',
        'parent_task_metadata': 'parent_task_metadata',
        'parent_node_metadata': 'parent_node_metadata',
        'retry_group': 'retry_group',
        'spec_node_id': 'spec_node_id',
        'node_name': 'node_name',
        'event_version': 'event_version',
        'is_parent': 'is_parent',
        'is_dynamic': 'is_dynamic',
        'deck_uri': 'deck_uri',
        'reported_at': 'reported_at'
    }

    def __init__(self, id=None, producer_id=None, phase=None, occurred_at=None, input_uri=None, input_data=None, output_uri=None, error=None, output_data=None, workflow_node_metadata=None, task_node_metadata=None, parent_task_metadata=None, parent_node_metadata=None, retry_group=None, spec_node_id=None, node_name=None, event_version=None, is_parent=None, is_dynamic=None, deck_uri=None, reported_at=None):  # noqa: E501
        """EventNodeExecutionEvent - a model defined in Swagger"""  # noqa: E501

        self._id = None
        self._producer_id = None
        self._phase = None
        self._occurred_at = None
        self._input_uri = None
        self._input_data = None
        self._output_uri = None
        self._error = None
        self._output_data = None
        self._workflow_node_metadata = None
        self._task_node_metadata = None
        self._parent_task_metadata = None
        self._parent_node_metadata = None
        self._retry_group = None
        self._spec_node_id = None
        self._node_name = None
        self._event_version = None
        self._is_parent = None
        self._is_dynamic = None
        self._deck_uri = None
        self._reported_at = None
        self.discriminator = None

        if id is not None:
            self.id = id
        if producer_id is not None:
            self.producer_id = producer_id
        if phase is not None:
            self.phase = phase
        if occurred_at is not None:
            self.occurred_at = occurred_at
        if input_uri is not None:
            self.input_uri = input_uri
        if input_data is not None:
            self.input_data = input_data
        if output_uri is not None:
            self.output_uri = output_uri
        if error is not None:
            self.error = error
        if output_data is not None:
            self.output_data = output_data
        if workflow_node_metadata is not None:
            self.workflow_node_metadata = workflow_node_metadata
        if task_node_metadata is not None:
            self.task_node_metadata = task_node_metadata
        if parent_task_metadata is not None:
            self.parent_task_metadata = parent_task_metadata
        if parent_node_metadata is not None:
            self.parent_node_metadata = parent_node_metadata
        if retry_group is not None:
            self.retry_group = retry_group
        if spec_node_id is not None:
            self.spec_node_id = spec_node_id
        if node_name is not None:
            self.node_name = node_name
        if event_version is not None:
            self.event_version = event_version
        if is_parent is not None:
            self.is_parent = is_parent
        if is_dynamic is not None:
            self.is_dynamic = is_dynamic
        if deck_uri is not None:
            self.deck_uri = deck_uri
        if reported_at is not None:
            self.reported_at = reported_at

    @property
    def id(self):
        """Gets the id of this EventNodeExecutionEvent.  # noqa: E501


        :return: The id of this EventNodeExecutionEvent.  # noqa: E501
        :rtype: CoreNodeExecutionIdentifier
        """
        return self._id

    @id.setter
    def id(self, id):
        """Sets the id of this EventNodeExecutionEvent.


        :param id: The id of this EventNodeExecutionEvent.  # noqa: E501
        :type: CoreNodeExecutionIdentifier
        """

        self._id = id

    @property
    def producer_id(self):
        """Gets the producer_id of this EventNodeExecutionEvent.  # noqa: E501


        :return: The producer_id of this EventNodeExecutionEvent.  # noqa: E501
        :rtype: str
        """
        return self._producer_id

    @producer_id.setter
    def producer_id(self, producer_id):
        """Sets the producer_id of this EventNodeExecutionEvent.


        :param producer_id: The producer_id of this EventNodeExecutionEvent.  # noqa: E501
        :type: str
        """

        self._producer_id = producer_id

    @property
    def phase(self):
        """Gets the phase of this EventNodeExecutionEvent.  # noqa: E501


        :return: The phase of this EventNodeExecutionEvent.  # noqa: E501
        :rtype: CoreNodeExecutionPhase
        """
        return self._phase

    @phase.setter
    def phase(self, phase):
        """Sets the phase of this EventNodeExecutionEvent.


        :param phase: The phase of this EventNodeExecutionEvent.  # noqa: E501
        :type: CoreNodeExecutionPhase
        """

        self._phase = phase

    @property
    def occurred_at(self):
        """Gets the occurred_at of this EventNodeExecutionEvent.  # noqa: E501

        This timestamp represents when the original event occurred, it is generated by the executor of the node.  # noqa: E501

        :return: The occurred_at of this EventNodeExecutionEvent.  # noqa: E501
        :rtype: datetime
        """
        return self._occurred_at

    @occurred_at.setter
    def occurred_at(self, occurred_at):
        """Sets the occurred_at of this EventNodeExecutionEvent.

        This timestamp represents when the original event occurred, it is generated by the executor of the node.  # noqa: E501

        :param occurred_at: The occurred_at of this EventNodeExecutionEvent.  # noqa: E501
        :type: datetime
        """

        self._occurred_at = occurred_at

    @property
    def input_uri(self):
        """Gets the input_uri of this EventNodeExecutionEvent.  # noqa: E501


        :return: The input_uri of this EventNodeExecutionEvent.  # noqa: E501
        :rtype: str
        """
        return self._input_uri

    @input_uri.setter
    def input_uri(self, input_uri):
        """Sets the input_uri of this EventNodeExecutionEvent.


        :param input_uri: The input_uri of this EventNodeExecutionEvent.  # noqa: E501
        :type: str
        """

        self._input_uri = input_uri

    @property
    def input_data(self):
        """Gets the input_data of this EventNodeExecutionEvent.  # noqa: E501

        Raw input data consumed by this node execution.  # noqa: E501

        :return: The input_data of this EventNodeExecutionEvent.  # noqa: E501
        :rtype: CoreLiteralMap
        """
        return self._input_data

    @input_data.setter
    def input_data(self, input_data):
        """Sets the input_data of this EventNodeExecutionEvent.

        Raw input data consumed by this node execution.  # noqa: E501

        :param input_data: The input_data of this EventNodeExecutionEvent.  # noqa: E501
        :type: CoreLiteralMap
        """

        self._input_data = input_data

    @property
    def output_uri(self):
        """Gets the output_uri of this EventNodeExecutionEvent.  # noqa: E501

        URL to the output of the execution, it encodes all the information including Cloud source provider. ie., s3://...  # noqa: E501

        :return: The output_uri of this EventNodeExecutionEvent.  # noqa: E501
        :rtype: str
        """
        return self._output_uri

    @output_uri.setter
    def output_uri(self, output_uri):
        """Sets the output_uri of this EventNodeExecutionEvent.

        URL to the output of the execution, it encodes all the information including Cloud source provider. ie., s3://...  # noqa: E501

        :param output_uri: The output_uri of this EventNodeExecutionEvent.  # noqa: E501
        :type: str
        """

        self._output_uri = output_uri

    @property
    def error(self):
        """Gets the error of this EventNodeExecutionEvent.  # noqa: E501


        :return: The error of this EventNodeExecutionEvent.  # noqa: E501
        :rtype: CoreExecutionError
        """
        return self._error

    @error.setter
    def error(self, error):
        """Sets the error of this EventNodeExecutionEvent.


        :param error: The error of this EventNodeExecutionEvent.  # noqa: E501
        :type: CoreExecutionError
        """

        self._error = error

    @property
    def output_data(self):
        """Gets the output_data of this EventNodeExecutionEvent.  # noqa: E501

        Raw output data produced by this node execution.  # noqa: E501

        :return: The output_data of this EventNodeExecutionEvent.  # noqa: E501
        :rtype: CoreLiteralMap
        """
        return self._output_data

    @output_data.setter
    def output_data(self, output_data):
        """Sets the output_data of this EventNodeExecutionEvent.

        Raw output data produced by this node execution.  # noqa: E501

        :param output_data: The output_data of this EventNodeExecutionEvent.  # noqa: E501
        :type: CoreLiteralMap
        """

        self._output_data = output_data

    @property
    def workflow_node_metadata(self):
        """Gets the workflow_node_metadata of this EventNodeExecutionEvent.  # noqa: E501


        :return: The workflow_node_metadata of this EventNodeExecutionEvent.  # noqa: E501
        :rtype: FlyteidleventWorkflowNodeMetadata
        """
        return self._workflow_node_metadata

    @workflow_node_metadata.setter
    def workflow_node_metadata(self, workflow_node_metadata):
        """Sets the workflow_node_metadata of this EventNodeExecutionEvent.


        :param workflow_node_metadata: The workflow_node_metadata of this EventNodeExecutionEvent.  # noqa: E501
        :type: FlyteidleventWorkflowNodeMetadata
        """

        self._workflow_node_metadata = workflow_node_metadata

    @property
    def task_node_metadata(self):
        """Gets the task_node_metadata of this EventNodeExecutionEvent.  # noqa: E501


        :return: The task_node_metadata of this EventNodeExecutionEvent.  # noqa: E501
        :rtype: FlyteidleventTaskNodeMetadata
        """
        return self._task_node_metadata

    @task_node_metadata.setter
    def task_node_metadata(self, task_node_metadata):
        """Sets the task_node_metadata of this EventNodeExecutionEvent.


        :param task_node_metadata: The task_node_metadata of this EventNodeExecutionEvent.  # noqa: E501
        :type: FlyteidleventTaskNodeMetadata
        """

        self._task_node_metadata = task_node_metadata

    @property
    def parent_task_metadata(self):
        """Gets the parent_task_metadata of this EventNodeExecutionEvent.  # noqa: E501

        [To be deprecated] Specifies which task (if any) launched this node.  # noqa: E501

        :return: The parent_task_metadata of this EventNodeExecutionEvent.  # noqa: E501
        :rtype: EventParentTaskExecutionMetadata
        """
        return self._parent_task_metadata

    @parent_task_metadata.setter
    def parent_task_metadata(self, parent_task_metadata):
        """Sets the parent_task_metadata of this EventNodeExecutionEvent.

        [To be deprecated] Specifies which task (if any) launched this node.  # noqa: E501

        :param parent_task_metadata: The parent_task_metadata of this EventNodeExecutionEvent.  # noqa: E501
        :type: EventParentTaskExecutionMetadata
        """

        self._parent_task_metadata = parent_task_metadata

    @property
    def parent_node_metadata(self):
        """Gets the parent_node_metadata of this EventNodeExecutionEvent.  # noqa: E501

        Specifies the parent node of the current node execution. Node executions at level zero will not have a parent node.  # noqa: E501

        :return: The parent_node_metadata of this EventNodeExecutionEvent.  # noqa: E501
        :rtype: EventParentNodeExecutionMetadata
        """
        return self._parent_node_metadata

    @parent_node_metadata.setter
    def parent_node_metadata(self, parent_node_metadata):
        """Sets the parent_node_metadata of this EventNodeExecutionEvent.

        Specifies the parent node of the current node execution. Node executions at level zero will not have a parent node.  # noqa: E501

        :param parent_node_metadata: The parent_node_metadata of this EventNodeExecutionEvent.  # noqa: E501
        :type: EventParentNodeExecutionMetadata
        """

        self._parent_node_metadata = parent_node_metadata

    @property
    def retry_group(self):
        """Gets the retry_group of this EventNodeExecutionEvent.  # noqa: E501


        :return: The retry_group of this EventNodeExecutionEvent.  # noqa: E501
        :rtype: str
        """
        return self._retry_group

    @retry_group.setter
    def retry_group(self, retry_group):
        """Sets the retry_group of this EventNodeExecutionEvent.


        :param retry_group: The retry_group of this EventNodeExecutionEvent.  # noqa: E501
        :type: str
        """

        self._retry_group = retry_group

    @property
    def spec_node_id(self):
        """Gets the spec_node_id of this EventNodeExecutionEvent.  # noqa: E501


        :return: The spec_node_id of this EventNodeExecutionEvent.  # noqa: E501
        :rtype: str
        """
        return self._spec_node_id

    @spec_node_id.setter
    def spec_node_id(self, spec_node_id):
        """Sets the spec_node_id of this EventNodeExecutionEvent.


        :param spec_node_id: The spec_node_id of this EventNodeExecutionEvent.  # noqa: E501
        :type: str
        """

        self._spec_node_id = spec_node_id

    @property
    def node_name(self):
        """Gets the node_name of this EventNodeExecutionEvent.  # noqa: E501


        :return: The node_name of this EventNodeExecutionEvent.  # noqa: E501
        :rtype: str
        """
        return self._node_name

    @node_name.setter
    def node_name(self, node_name):
        """Sets the node_name of this EventNodeExecutionEvent.


        :param node_name: The node_name of this EventNodeExecutionEvent.  # noqa: E501
        :type: str
        """

        self._node_name = node_name

    @property
    def event_version(self):
        """Gets the event_version of this EventNodeExecutionEvent.  # noqa: E501


        :return: The event_version of this EventNodeExecutionEvent.  # noqa: E501
        :rtype: int
        """
        return self._event_version

    @event_version.setter
    def event_version(self, event_version):
        """Sets the event_version of this EventNodeExecutionEvent.


        :param event_version: The event_version of this EventNodeExecutionEvent.  # noqa: E501
        :type: int
        """

        self._event_version = event_version

    @property
    def is_parent(self):
        """Gets the is_parent of this EventNodeExecutionEvent.  # noqa: E501

        Whether this node launched a subworkflow.  # noqa: E501

        :return: The is_parent of this EventNodeExecutionEvent.  # noqa: E501
        :rtype: bool
        """
        return self._is_parent

    @is_parent.setter
    def is_parent(self, is_parent):
        """Sets the is_parent of this EventNodeExecutionEvent.

        Whether this node launched a subworkflow.  # noqa: E501

        :param is_parent: The is_parent of this EventNodeExecutionEvent.  # noqa: E501
        :type: bool
        """

        self._is_parent = is_parent

    @property
    def is_dynamic(self):
        """Gets the is_dynamic of this EventNodeExecutionEvent.  # noqa: E501

        Whether this node yielded a dynamic workflow.  # noqa: E501

        :return: The is_dynamic of this EventNodeExecutionEvent.  # noqa: E501
        :rtype: bool
        """
        return self._is_dynamic

    @is_dynamic.setter
    def is_dynamic(self, is_dynamic):
        """Sets the is_dynamic of this EventNodeExecutionEvent.

        Whether this node yielded a dynamic workflow.  # noqa: E501

        :param is_dynamic: The is_dynamic of this EventNodeExecutionEvent.  # noqa: E501
        :type: bool
        """

        self._is_dynamic = is_dynamic

    @property
    def deck_uri(self):
        """Gets the deck_uri of this EventNodeExecutionEvent.  # noqa: E501


        :return: The deck_uri of this EventNodeExecutionEvent.  # noqa: E501
        :rtype: str
        """
        return self._deck_uri

    @deck_uri.setter
    def deck_uri(self, deck_uri):
        """Sets the deck_uri of this EventNodeExecutionEvent.


        :param deck_uri: The deck_uri of this EventNodeExecutionEvent.  # noqa: E501
        :type: str
        """

        self._deck_uri = deck_uri

    @property
    def reported_at(self):
        """Gets the reported_at of this EventNodeExecutionEvent.  # noqa: E501

        This timestamp represents the instant when the event was reported by the executing framework. For example, when first processing a node the `occurred_at` timestamp should be the instant propeller makes progress, so when literal inputs are initially copied. The event however will not be sent until after the copy completes. Extracting both of these timestamps facilitates a more accurate portrayal of the evaluation time-series.  # noqa: E501

        :return: The reported_at of this EventNodeExecutionEvent.  # noqa: E501
        :rtype: datetime
        """
        return self._reported_at

    @reported_at.setter
    def reported_at(self, reported_at):
        """Sets the reported_at of this EventNodeExecutionEvent.

        This timestamp represents the instant when the event was reported by the executing framework. For example, when first processing a node the `occurred_at` timestamp should be the instant propeller makes progress, so when literal inputs are initially copied. The event however will not be sent until after the copy completes. Extracting both of these timestamps facilitates a more accurate portrayal of the evaluation time-series.  # noqa: E501

        :param reported_at: The reported_at of this EventNodeExecutionEvent.  # noqa: E501
        :type: datetime
        """

        self._reported_at = reported_at

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value
        if issubclass(EventNodeExecutionEvent, dict):
            for key, value in self.items():
                result[key] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, EventNodeExecutionEvent):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
