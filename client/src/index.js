


const {EventLoggerClient} =  requeire('./generated/event_grpc_pb.js');
const { LoadEventsRequest,
        Filter,
        LoadEventsResponse,
        EventsList,
        Event
        } = require('../generated/event_pb')
const {EventsLoggerApp} = require('./eventsloggerapp.js')

var eventLoggerService = new EventLoggerClient("http://localhost:8080", null, null)

var eventsLoggerApp = new EventsLoggerApp(
    eventLoggerService, {
        LoadEventsRequest:LoadEventsRequest,
        Filter:Filter,
        LoadEventsResponse:LoadEventsResponse,
        EventsList:EventsList,
        Event:Event
    },
    {
        checkGrpcStatusCode: function(status) {
            if (status.code != grpc.web.StatusCode.OK) {
                console.log('Error code: '+status.code+' "'+
                    status.details+'"');
            }
        }
    }
);

eventLoggerService.load()
