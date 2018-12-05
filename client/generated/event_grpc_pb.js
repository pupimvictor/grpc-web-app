/**
 * @fileoverview gRPC-Web generated client stub for server
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.server = require('./event_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.server.EventLoggerClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.server.EventLoggerPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!proto.server.EventLoggerClient} The delegate callback based client
   */
  this.delegateClient_ = new proto.server.EventLoggerClient(
      hostname, credentials, options);

};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.server.LoadEventsRequest,
 *   !proto.server.LoadEventsResponse>}
 */
const methodInfo_EventLogger_LoadEvents = new grpc.web.AbstractClientBase.MethodInfo(
  proto.server.LoadEventsResponse,
  /** @param {!proto.server.LoadEventsRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.server.LoadEventsResponse.deserializeBinary
);


/**
 * @param {!proto.server.LoadEventsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.server.LoadEventsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.server.LoadEventsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.server.EventLoggerClient.prototype.loadEvents =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/server.EventLogger/LoadEvents',
      request,
      metadata || {},
      methodInfo_EventLogger_LoadEvents,
      callback);
};


/**
 * @param {!proto.server.LoadEventsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.server.LoadEventsResponse>}
 *     The XHR Node Readable Stream
 */
proto.server.EventLoggerPromiseClient.prototype.loadEvents =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.loadEvents(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.server.StreamEventsRequest,
 *   !proto.server.StreamEventsResponse>}
 */
const methodInfo_EventLogger_StreamEvents = new grpc.web.AbstractClientBase.MethodInfo(
  proto.server.StreamEventsResponse,
  /** @param {!proto.server.StreamEventsRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.server.StreamEventsResponse.deserializeBinary
);


/**
 * @param {!proto.server.StreamEventsRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.server.StreamEventsResponse>}
 *     The XHR Node Readable Stream
 */
proto.server.EventLoggerClient.prototype.streamEvents =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/server.EventLogger/StreamEvents',
      request,
      metadata,
      methodInfo_EventLogger_StreamEvents);
};


/**
 * @param {!proto.server.StreamEventsRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.server.StreamEventsResponse>}
 *     The XHR Node Readable Stream
 */
proto.server.EventLoggerPromiseClient.prototype.streamEvents =
    function(request, metadata) {
  return this.delegateClient_.client_.serverStreaming(this.delegateClient_.hostname_ +
      '/server.EventLogger/StreamEvents',
      request,
      metadata,
      methodInfo_EventLogger_StreamEvents);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.server.StreamId,
 *   !proto.server.Void>}
 */
const methodInfo_EventLogger_StopStreaming = new grpc.web.AbstractClientBase.MethodInfo(
  proto.server.Void,
  /** @param {!proto.server.StreamId} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.server.Void.deserializeBinary
);


/**
 * @param {!proto.server.StreamId} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.server.Void)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.server.Void>|undefined}
 *     The XHR Node Readable Stream
 */
proto.server.EventLoggerClient.prototype.stopStreaming =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/server.EventLogger/StopStreaming',
      request,
      metadata || {},
      methodInfo_EventLogger_StopStreaming,
      callback);
};


/**
 * @param {!proto.server.StreamId} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.server.Void>}
 *     The XHR Node Readable Stream
 */
proto.server.EventLoggerPromiseClient.prototype.stopStreaming =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.stopStreaming(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


module.exports = proto.server;

