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

/**
 * The Repository model module.
 * @module model/Repository
 * @version 1.0.0
 */
class Repository {
    /**
     * Constructs a new <code>Repository</code>.
     * @alias module:model/Repository
     */
    constructor() { 
        
        Repository.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>Repository</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/Repository} obj Optional instance to populate.
     * @return {module:model/Repository} The populated <code>Repository</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new Repository();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('private')) {
                obj['private'] = ApiClient.convertToType(data['private'], 'Boolean');
            }
            if (data.hasOwnProperty('disabled')) {
                obj['disabled'] = ApiClient.convertToType(data['disabled'], 'Boolean');
            }
            if (data.hasOwnProperty('auto_created')) {
                obj['auto_created'] = ApiClient.convertToType(data['auto_created'], 'Boolean');
            }
            if (data.hasOwnProperty('github')) {
                obj['github'] = ApiClient.convertToType(data['github'], Object);
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
Repository.prototype['id'] = undefined;

/**
 * @member {String} name
 */
Repository.prototype['name'] = undefined;

/**
 * @member {Boolean} private
 */
Repository.prototype['private'] = undefined;

/**
 * @member {Boolean} disabled
 */
Repository.prototype['disabled'] = undefined;

/**
 * @member {Boolean} auto_created
 */
Repository.prototype['auto_created'] = undefined;

/**
 * @member {Object} github
 */
Repository.prototype['github'] = undefined;






export default Repository;

