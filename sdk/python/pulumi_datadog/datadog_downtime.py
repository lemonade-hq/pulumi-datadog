# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class DatadogDowntime(pulumi.CustomResource):
    active: pulumi.Output[bool]
    """
    A flag indicating if the downtime is active now.
    """
    disabled: pulumi.Output[bool]
    """
    A flag indicating if the downtime was disabled.
    """
    end: pulumi.Output[float]
    """
    POSIX timestamp to end the downtime.
    """
    end_date: pulumi.Output[str]
    """
    String representing date and time to end the downtime in RFC3339 format.
    """
    message: pulumi.Output[str]
    """
    A message to include with notifications for this downtime.
    """
    monitor_id: pulumi.Output[float]
    """
    Reference to which monitor this downtime is applied. When scheduling downtime for a given monitor, datadog changes `silenced` property of the monitor  to match the `end` POSIX timestamp.
    """
    recurrence: pulumi.Output[dict]
    """
    A dictionary to configure the downtime to be recurring.
    """
    scopes: pulumi.Output[list]
    """
    A list of items to apply the downtime to, e.g. host:X
    """
    start: pulumi.Output[float]
    """
    POSIX timestamp to start the downtime.
    """
    start_date: pulumi.Output[str]
    """
    String representing date and time to start the downtime in RFC3339 format.
    """
    def __init__(__self__, resource_name, opts=None, active=None, disabled=None, end=None, end_date=None, message=None, monitor_id=None, recurrence=None, scopes=None, start=None, start_date=None, __name__=None, __opts__=None):
        """
        Provides a Datadog downtime resource. This can be used to create and manage Datadog downtimes.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] active: A flag indicating if the downtime is active now.
        :param pulumi.Input[bool] disabled: A flag indicating if the downtime was disabled.
        :param pulumi.Input[float] end: POSIX timestamp to end the downtime.
        :param pulumi.Input[str] end_date: String representing date and time to end the downtime in RFC3339 format.
        :param pulumi.Input[str] message: A message to include with notifications for this downtime.
        :param pulumi.Input[float] monitor_id: Reference to which monitor this downtime is applied. When scheduling downtime for a given monitor, datadog changes `silenced` property of the monitor  to match the `end` POSIX timestamp.
        :param pulumi.Input[dict] recurrence: A dictionary to configure the downtime to be recurring.
        :param pulumi.Input[list] scopes: A list of items to apply the downtime to, e.g. host:X
        :param pulumi.Input[float] start: POSIX timestamp to start the downtime.
        :param pulumi.Input[str] start_date: String representing date and time to start the downtime in RFC3339 format.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['active'] = active

        __props__['disabled'] = disabled

        __props__['end'] = end

        __props__['end_date'] = end_date

        __props__['message'] = message

        __props__['monitor_id'] = monitor_id

        __props__['recurrence'] = recurrence

        if scopes is None:
            raise TypeError("Missing required property 'scopes'")
        __props__['scopes'] = scopes

        __props__['start'] = start

        __props__['start_date'] = start_date

        super(DatadogDowntime, __self__).__init__(
            'datadog:index/datadogDowntime:DatadogDowntime',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

