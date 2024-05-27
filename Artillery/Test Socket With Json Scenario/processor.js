'use strict';

const scenarios = require("./scenarios.json");

module.exports.handleBidding = function (context, events, done) {
  context.vars.greetingMessage = scenarios.message;
  return done();
}