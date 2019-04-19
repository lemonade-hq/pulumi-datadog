// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class DatadogMetricMetadata extends pulumi.CustomResource {
    /**
     * Get an existing DatadogMetricMetadata resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatadogMetricMetadataState, opts?: pulumi.CustomResourceOptions): DatadogMetricMetadata {
        return new DatadogMetricMetadata(name, <any>state, { ...opts, id: id });
    }

    public readonly description: pulumi.Output<string | undefined>;
    public readonly metric: pulumi.Output<string>;
    public readonly perUnit: pulumi.Output<string | undefined>;
    public readonly shortName: pulumi.Output<string | undefined>;
    public readonly statsdInterval: pulumi.Output<number | undefined>;
    public readonly type: pulumi.Output<string | undefined>;
    public readonly unit: pulumi.Output<string | undefined>;

    /**
     * Create a DatadogMetricMetadata resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatadogMetricMetadataArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatadogMetricMetadataArgs | DatadogMetricMetadataState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: DatadogMetricMetadataState = argsOrState as DatadogMetricMetadataState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["metric"] = state ? state.metric : undefined;
            inputs["perUnit"] = state ? state.perUnit : undefined;
            inputs["shortName"] = state ? state.shortName : undefined;
            inputs["statsdInterval"] = state ? state.statsdInterval : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["unit"] = state ? state.unit : undefined;
        } else {
            const args = argsOrState as DatadogMetricMetadataArgs | undefined;
            if (!args || args.metric === undefined) {
                throw new Error("Missing required property 'metric'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["metric"] = args ? args.metric : undefined;
            inputs["perUnit"] = args ? args.perUnit : undefined;
            inputs["shortName"] = args ? args.shortName : undefined;
            inputs["statsdInterval"] = args ? args.statsdInterval : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["unit"] = args ? args.unit : undefined;
        }
        super("datadog:index/datadogMetricMetadata:DatadogMetricMetadata", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DatadogMetricMetadata resources.
 */
export interface DatadogMetricMetadataState {
    readonly description?: pulumi.Input<string>;
    readonly metric?: pulumi.Input<string>;
    readonly perUnit?: pulumi.Input<string>;
    readonly shortName?: pulumi.Input<string>;
    readonly statsdInterval?: pulumi.Input<number>;
    readonly type?: pulumi.Input<string>;
    readonly unit?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DatadogMetricMetadata resource.
 */
export interface DatadogMetricMetadataArgs {
    readonly description?: pulumi.Input<string>;
    readonly metric: pulumi.Input<string>;
    readonly perUnit?: pulumi.Input<string>;
    readonly shortName?: pulumi.Input<string>;
    readonly statsdInterval?: pulumi.Input<number>;
    readonly type?: pulumi.Input<string>;
    readonly unit?: pulumi.Input<string>;
}
