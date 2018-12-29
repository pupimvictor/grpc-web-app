


const eventsapp = {}

eventsapp.App = function(eventsServer, msgs, handlers) {
    this.eventsServer = eventsServer;
    this.msgs = msgs;
    this.handlers = handlers;
}

eventsapp.App.prototype.loadEvents = function(f){
    var self = this;
    var filter = new this.msgs.Filter();
    filter.setMsg(f.msg);
    filter.setBasedate(f.baseDate);
    filter.setSeverityid(f.severity);
    filter.setSystem(f.system);

    var loadEventsRequest = new this.msgs.LoadEventsRequest();
    loadEventsRequest.setFilter(filter);

    var call = this.eventsServer.loadEvents(loadEventsRequest, {"header-test": "test-val"},
        function(err, response){
            if (err) {
                console.log(err)
            } else {
                console.log(response)
            }
    });
    call.on('status', function(status){
       self.handlers.checkGrpcStatusCode(status);
       if (status.metadata) {
           console.log("status meta: " + status.metadata)
       }
    });
};

eventsapp.App.prototype.load = function(){
    console.log("load app")
    this.loadEvents(({msg: "", baseDate: 1544049929, severity: 1, system: "A"}))
};