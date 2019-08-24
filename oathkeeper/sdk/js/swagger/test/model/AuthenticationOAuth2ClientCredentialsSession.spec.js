/**
 * ORY Oathkeeper
 * ORY Oathkeeper is a reverse proxy that checks the HTTP Authorization for validity against a set of rules. This service uses Hydra to validate access tokens and policies.
 *
 * OpenAPI spec version: Latest
 * Contact: hi@ory.am
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.2.3
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', '../../src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require('../../src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.OryOathkeeper);
  }
}(this, function(expect, OryOathkeeper) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new OryOathkeeper.AuthenticationOAuth2ClientCredentialsSession();
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

  describe('AuthenticationOAuth2ClientCredentialsSession', function() {
    it('should create an instance of AuthenticationOAuth2ClientCredentialsSession', function() {
      // uncomment below and update the code to test AuthenticationOAuth2ClientCredentialsSession
      //var instane = new OryOathkeeper.AuthenticationOAuth2ClientCredentialsSession();
      //expect(instance).to.be.a(OryOathkeeper.AuthenticationOAuth2ClientCredentialsSession);
    });

    it('should have the property allowed (base name: "allowed")', function() {
      // uncomment below and update the code to test the property allowed
      //var instane = new OryOathkeeper.AuthenticationOAuth2ClientCredentialsSession();
      //expect(instance).to.be();
    });

    it('should have the property sub (base name: "sub")', function() {
      // uncomment below and update the code to test the property sub
      //var instane = new OryOathkeeper.AuthenticationOAuth2ClientCredentialsSession();
      //expect(instance).to.be();
    });

  });

}));
