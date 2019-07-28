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
    instance = new Uisvc.DefaultApi();
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

  describe('DefaultApi', function() {
    describe('cancelRunIdPost', function() {
      it('should call cancelRunIdPost successfully', function(done) {
        //uncomment below and update the code to test cancelRunIdPost
        //instance.cancelRunIdPost(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('capabilitiesUsernameCapabilityDelete', function() {
      it('should call capabilitiesUsernameCapabilityDelete successfully', function(done) {
        //uncomment below and update the code to test capabilitiesUsernameCapabilityDelete
        //instance.capabilitiesUsernameCapabilityDelete(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('capabilitiesUsernameCapabilityPost', function() {
      it('should call capabilitiesUsernameCapabilityPost successfully', function(done) {
        //uncomment below and update the code to test capabilitiesUsernameCapabilityPost
        //instance.capabilitiesUsernameCapabilityPost(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('errorsGet', function() {
      it('should call errorsGet successfully', function(done) {
        //uncomment below and update the code to test errorsGet
        //instance.errorsGet(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('logAttachIdGet', function() {
      it('should call logAttachIdGet successfully', function(done) {
        //uncomment below and update the code to test logAttachIdGet
        //instance.logAttachIdGet(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('loggedinGet', function() {
      it('should call loggedinGet successfully', function(done) {
        //uncomment below and update the code to test loggedinGet
        //instance.loggedinGet(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('loginGet', function() {
      it('should call loginGet successfully', function(done) {
        //uncomment below and update the code to test loginGet
        //instance.loginGet(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('loginUpgradeGet', function() {
      it('should call loginUpgradeGet successfully', function(done) {
        //uncomment below and update the code to test loginUpgradeGet
        //instance.loginUpgradeGet(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('logoutGet', function() {
      it('should call logoutGet successfully', function(done) {
        //uncomment below and update the code to test logoutGet
        //instance.logoutGet(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('repositoriesCiAddOwnerRepoGet', function() {
      it('should call repositoriesCiAddOwnerRepoGet successfully', function(done) {
        //uncomment below and update the code to test repositoriesCiAddOwnerRepoGet
        //instance.repositoriesCiAddOwnerRepoGet(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('repositoriesCiDelOwnerRepoGet', function() {
      it('should call repositoriesCiDelOwnerRepoGet successfully', function(done) {
        //uncomment below and update the code to test repositoriesCiDelOwnerRepoGet
        //instance.repositoriesCiDelOwnerRepoGet(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('repositoriesMyGet', function() {
      it('should call repositoriesMyGet successfully', function(done) {
        //uncomment below and update the code to test repositoriesMyGet
        //instance.repositoriesMyGet(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('repositoriesScanGet', function() {
      it('should call repositoriesScanGet successfully', function(done) {
        //uncomment below and update the code to test repositoriesScanGet
        //instance.repositoriesScanGet(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('repositoriesSubAddOwnerRepoGet', function() {
      it('should call repositoriesSubAddOwnerRepoGet successfully', function(done) {
        //uncomment below and update the code to test repositoriesSubAddOwnerRepoGet
        //instance.repositoriesSubAddOwnerRepoGet(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('repositoriesSubDelOwnerRepoGet', function() {
      it('should call repositoriesSubDelOwnerRepoGet successfully', function(done) {
        //uncomment below and update the code to test repositoriesSubDelOwnerRepoGet
        //instance.repositoriesSubDelOwnerRepoGet(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('repositoriesSubscribedGet', function() {
      it('should call repositoriesSubscribedGet successfully', function(done) {
        //uncomment below and update the code to test repositoriesSubscribedGet
        //instance.repositoriesSubscribedGet(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('repositoriesVisibleGet', function() {
      it('should call repositoriesVisibleGet successfully', function(done) {
        //uncomment below and update the code to test repositoriesVisibleGet
        //instance.repositoriesVisibleGet(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('runRunIdGet', function() {
      it('should call runRunIdGet successfully', function(done) {
        //uncomment below and update the code to test runRunIdGet
        //instance.runRunIdGet(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('runsCountGet', function() {
      it('should call runsCountGet successfully', function(done) {
        //uncomment below and update the code to test runsCountGet
        //instance.runsCountGet(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('runsGet', function() {
      it('should call runsGet successfully', function(done) {
        //uncomment below and update the code to test runsGet
        //instance.runsGet(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('submissionIdGet', function() {
      it('should call submissionIdGet successfully', function(done) {
        //uncomment below and update the code to test submissionIdGet
        //instance.submissionIdGet(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('submissionIdTasksGet', function() {
      it('should call submissionIdTasksGet successfully', function(done) {
        //uncomment below and update the code to test submissionIdTasksGet
        //instance.submissionIdTasksGet(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('submissionsGet', function() {
      it('should call submissionsGet successfully', function(done) {
        //uncomment below and update the code to test submissionsGet
        //instance.submissionsGet(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('submitGet', function() {
      it('should call submitGet successfully', function(done) {
        //uncomment below and update the code to test submitGet
        //instance.submitGet(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('tasksCancelIdPost', function() {
      it('should call tasksCancelIdPost successfully', function(done) {
        //uncomment below and update the code to test tasksCancelIdPost
        //instance.tasksCancelIdPost(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('tasksCountGet', function() {
      it('should call tasksCountGet successfully', function(done) {
        //uncomment below and update the code to test tasksCountGet
        //instance.tasksCountGet(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('tasksGet', function() {
      it('should call tasksGet successfully', function(done) {
        //uncomment below and update the code to test tasksGet
        //instance.tasksGet(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('tasksRunsIdCountGet', function() {
      it('should call tasksRunsIdCountGet successfully', function(done) {
        //uncomment below and update the code to test tasksRunsIdCountGet
        //instance.tasksRunsIdCountGet(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('tasksRunsIdGet', function() {
      it('should call tasksRunsIdGet successfully', function(done) {
        //uncomment below and update the code to test tasksRunsIdGet
        //instance.tasksRunsIdGet(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('tasksSubscribedGet', function() {
      it('should call tasksSubscribedGet successfully', function(done) {
        //uncomment below and update the code to test tasksSubscribedGet
        //instance.tasksSubscribedGet(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('tokenDelete', function() {
      it('should call tokenDelete successfully', function(done) {
        //uncomment below and update the code to test tokenDelete
        //instance.tokenDelete(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('tokenGet', function() {
      it('should call tokenGet successfully', function(done) {
        //uncomment below and update the code to test tokenGet
        //instance.tokenGet(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('userPropertiesGet', function() {
      it('should call userPropertiesGet successfully', function(done) {
        //uncomment below and update the code to test userPropertiesGet
        //instance.userPropertiesGet(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
  });

}));
