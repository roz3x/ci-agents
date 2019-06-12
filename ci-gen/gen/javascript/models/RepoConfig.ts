// tslint:disable
/**
 * uisvc
 * API for the user interface service; the service that is directly responsible for presenting data to users. This service typically runs at the border, and leverages session cookies or authentication tokens that we generate for users. It also is responsible for handling the act of oauth and user creation through its login hooks. uisvc typically talks to the datasvc and other services to accomplish its goal, it does not save anything locally or carry state. 
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface RepoConfig
 */
export interface RepoConfig {
    /**
     * 
     * @type {string}
     * @memberof RepoConfig
     */
    workdir?: string;
    /**
     * 
     * @type {string}
     * @memberof RepoConfig
     */
    queue?: string;
    /**
     * 
     * @type {boolean}
     * @memberof RepoConfig
     */
    overrideQueue?: boolean;
    /**
     * 
     * @type {number}
     * @memberof RepoConfig
     */
    globalTimeout?: number;
    /**
     * 
     * @type {boolean}
     * @memberof RepoConfig
     */
    overrideTimeout?: boolean;
}

export function RepoConfigFromJSON(json: any): RepoConfig {
    return {
        'workdir': !exists(json, 'workdir') ? undefined : json['workdir'],
        'queue': !exists(json, 'queue') ? undefined : json['queue'],
        'overrideQueue': !exists(json, 'override_queue') ? undefined : json['override_queue'],
        'globalTimeout': !exists(json, 'global_timeout') ? undefined : json['global_timeout'],
        'overrideTimeout': !exists(json, 'override_timeout') ? undefined : json['override_timeout'],
    };
}

export function RepoConfigToJSON(value?: RepoConfig): any {
    if (value === undefined) {
        return undefined;
    }
    return {
        'workdir': value.workdir,
        'queue': value.queue,
        'override_queue': value.overrideQueue,
        'global_timeout': value.globalTimeout,
        'override_timeout': value.overrideTimeout,
    };
}


