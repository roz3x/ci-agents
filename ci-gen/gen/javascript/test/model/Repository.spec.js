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

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', process.cwd()+'/src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require(process.cwd()+'/src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.Uisvc);
  }
}(this, function(expect, Uisvc) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new Uisvc.Repository();
  });

  var getProperty = function(object, getter, property) {
    // Use getter method if present; otherwise, get the property directly.
    if (typeof object[getter] === 'function')
      return object[getter]();
    else
      return object[property];
  }

  var setProperty = function(object, setter, property, value) {
    // Use setter method if present; otherwise, set the property directly.
    if (typeof object[setter] === 'function')
      object[setter](value);
    else
      object[property] = value;
  }

  describe('Repository', function() {
    it('should create an instance of Repository', function() {
      // uncomment below and update the code to test Repository
      //var instane = new Uisvc.Repository();
      //expect(instance).to.be.a(Uisvc.Repository);
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instane = new Uisvc.Repository();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instane = new Uisvc.Repository();
      //expect(instance).to.be();
    });

    it('should have the property _private (base name: "private")', function() {
      // uncomment below and update the code to test the property _private
      //var instane = new Uisvc.Repository();
      //expect(instance).to.be();
    });

    it('should have the property disabled (base name: "disabled")', function() {
      // uncomment below and update the code to test the property disabled
      //var instane = new Uisvc.Repository();
      //expect(instance).to.be();
    });

    it('should have the property autoCreated (base name: "auto_created")', function() {
      // uncomment below and update the code to test the property autoCreated
      //var instane = new Uisvc.Repository();
      //expect(instance).to.be();
    });

    it('should have the property github (base name: "github")', function() {
      // uncomment below and update the code to test the property github
      //var instane = new Uisvc.Repository();
      //expect(instance).to.be();
    });

  });

}));
