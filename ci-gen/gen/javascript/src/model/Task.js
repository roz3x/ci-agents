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
 *
 */

import ApiClient from '../ApiClient';
import ModelSubmission from './ModelSubmission';
import Ref from './Ref';
import Repository from './Repository';
import TaskSettings from './TaskSettings';

/**
 * The Task model module.
 * @module model/Task
 * @version 1.0.0
 */
class Task {
    /**
     * Constructs a new <code>Task</code>.
     * @alias module:model/Task
     */
    constructor() { 
        
        Task.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>Task</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/Task} obj Optional instance to populate.
     * @return {module:model/Task} The populated <code>Task</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new Task();

            if (data.hasOwnProperty('path')) {
                obj['path'] = ApiClient.convertToType(data['path'], 'String');
            }
            if (data.hasOwnProperty('settings')) {
                obj['settings'] = TaskSettings.constructFromObject(data['settings']);
            }
            if (data.hasOwnProperty('parent')) {
                obj['parent'] = Repository.constructFromObject(data['parent']);
            }
            if (data.hasOwnProperty('ref')) {
                obj['ref'] = Ref.constructFromObject(data['ref']);
            }
            if (data.hasOwnProperty('base_sha')) {
                obj['base_sha'] = ApiClient.convertToType(data['base_sha'], 'String');
            }
            if (data.hasOwnProperty('canceled')) {
                obj['canceled'] = ApiClient.convertToType(data['canceled'], 'Boolean');
            }
            if (data.hasOwnProperty('finished_at')) {
                obj['finished_at'] = ApiClient.convertToType(data['finished_at'], 'Date');
            }
            if (data.hasOwnProperty('started_at')) {
                obj['started_at'] = ApiClient.convertToType(data['started_at'], 'Date');
            }
            if (data.hasOwnProperty('created_at')) {
                obj['created_at'] = ApiClient.convertToType(data['created_at'], 'Date');
            }
            if (data.hasOwnProperty('status')) {
                obj['status'] = ApiClient.convertToType(data['status'], 'Boolean');
            }
            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('runs')) {
                obj['runs'] = ApiClient.convertToType(data['runs'], 'Number');
            }
            if (data.hasOwnProperty('submission')) {
                obj['submission'] = ModelSubmission.constructFromObject(data['submission']);
            }
        }
        return obj;
    }


}

/**
 * @member {String} path
 */
Task.prototype['path'] = undefined;

/**
 * @member {module:model/TaskSettings} settings
 */
Task.prototype['settings'] = undefined;

/**
 * @member {module:model/Repository} parent
 */
Task.prototype['parent'] = undefined;

/**
 * @member {module:model/Ref} ref
 */
Task.prototype['ref'] = undefined;

/**
 * @member {String} base_sha
 */
Task.prototype['base_sha'] = undefined;

/**
 * @member {Boolean} canceled
 */
Task.prototype['canceled'] = undefined;

/**
 * @member {Date} finished_at
 */
Task.prototype['finished_at'] = undefined;

/**
 * @member {Date} started_at
 */
Task.prototype['started_at'] = undefined;

/**
 * @member {Date} created_at
 */
Task.prototype['created_at'] = undefined;

/**
 * @member {Boolean} status
 */
Task.prototype['status'] = undefined;

/**
 * @member {Number} id
 */
Task.prototype['id'] = undefined;

/**
 * @member {Number} runs
 */
Task.prototype['runs'] = undefined;

/**
 * @member {module:model/ModelSubmission} submission
 */
Task.prototype['submission'] = undefined;






export default Task;

