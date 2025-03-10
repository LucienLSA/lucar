// Common aliases
const $util = $protobuf.util;

// Exported root namespace
const $root = $protobuf.roots["default"] || ($protobuf.roots["default"] = {});

export const lucar = $root.lucar = (() => {

    /**
     * Namespace lucar.
     * @exports lucar
     * @namespace
     */
    const lucar = {};

    lucar.Location = (function() {

        /**
         * Properties of a Location.
         * @memberof lucar
         * @interface ILocation
         * @property {number|null} [latitude] Location latitude
         * @property {number|null} [longitude] Location longitude
         */

        /**
         * Constructs a new Location.
         * @memberof lucar
         * @classdesc Represents a Location.
         * @implements ILocation
         * @constructor
         * @param {lucar.ILocation=} [properties] Properties to set
         */
        function Location(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Location latitude.
         * @member {number} latitude
         * @memberof lucar.Location
         * @instance
         */
        Location.prototype.latitude = 0;

        /**
         * Location longitude.
         * @member {number} longitude
         * @memberof lucar.Location
         * @instance
         */
        Location.prototype.longitude = 0;

        /**
         * Creates a Location message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof lucar.Location
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {lucar.Location} Location
         */
        Location.fromObject = function fromObject(object) {
            if (object instanceof $root.lucar.Location)
                return object;
            let message = new $root.lucar.Location();
            if (object.latitude != null)
                message.latitude = Number(object.latitude);
            if (object.longitude != null)
                message.longitude = Number(object.longitude);
            return message;
        };

        /**
         * Creates a plain object from a Location message. Also converts values to other types if specified.
         * @function toObject
         * @memberof lucar.Location
         * @static
         * @param {lucar.Location} message Location
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Location.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults) {
                object.latitude = 0;
                object.longitude = 0;
            }
            if (message.latitude != null && message.hasOwnProperty("latitude"))
                object.latitude = options.json && !isFinite(message.latitude) ? String(message.latitude) : message.latitude;
            if (message.longitude != null && message.hasOwnProperty("longitude"))
                object.longitude = options.json && !isFinite(message.longitude) ? String(message.longitude) : message.longitude;
            return object;
        };

        /**
         * Converts this Location to JSON.
         * @function toJSON
         * @memberof lucar.Location
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        Location.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for Location
         * @function getTypeUrl
         * @memberof lucar.Location
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        Location.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
            if (typeUrlPrefix === undefined) {
                typeUrlPrefix = "type.googleapis.com";
            }
            return typeUrlPrefix + "/lucar.Location";
        };

        return Location;
    })();

    /**
     * TripStatus enum.
     * @name lucar.TripStatus
     * @enum {number}
     * @property {number} TS_NOT_SPECIFIED=0 TS_NOT_SPECIFIED value
     * @property {number} NOT_STARTED=1 NOT_STARTED value
     * @property {number} IN_PROGRESS=2 IN_PROGRESS value
     * @property {number} FINISHED=3 FINISHED value
     * @property {number} PAID=4 PAID value
     */
    lucar.TripStatus = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "TS_NOT_SPECIFIED"] = 0;
        values[valuesById[1] = "NOT_STARTED"] = 1;
        values[valuesById[2] = "IN_PROGRESS"] = 2;
        values[valuesById[3] = "FINISHED"] = 3;
        values[valuesById[4] = "PAID"] = 4;
        return values;
    })();

    lucar.Trip = (function() {

        /**
         * Properties of a Trip.
         * @memberof lucar
         * @interface ITrip
         * @property {string|null} [start] Trip start
         * @property {string|null} [end] Trip end
         * @property {lucar.ILocation|null} [startPos] Trip startPos
         * @property {Array.<lucar.ILocation>|null} [pathLocations] Trip pathLocations
         * @property {number|null} [durationSec] Trip durationSec
         * @property {lucar.ILocation|null} [endPos] Trip endPos
         * @property {number|null} [feeCent] Trip feeCent
         * @property {lucar.TripStatus|null} [status] Trip status
         */

        /**
         * Constructs a new Trip.
         * @memberof lucar
         * @classdesc Represents a Trip.
         * @implements ITrip
         * @constructor
         * @param {lucar.ITrip=} [properties] Properties to set
         */
        function Trip(properties) {
            this.pathLocations = [];
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Trip start.
         * @member {string} start
         * @memberof lucar.Trip
         * @instance
         */
        Trip.prototype.start = "";

        /**
         * Trip end.
         * @member {string} end
         * @memberof lucar.Trip
         * @instance
         */
        Trip.prototype.end = "";

        /**
         * Trip startPos.
         * @member {lucar.ILocation|null|undefined} startPos
         * @memberof lucar.Trip
         * @instance
         */
        Trip.prototype.startPos = null;

        /**
         * Trip pathLocations.
         * @member {Array.<lucar.ILocation>} pathLocations
         * @memberof lucar.Trip
         * @instance
         */
        Trip.prototype.pathLocations = $util.emptyArray;

        /**
         * Trip durationSec.
         * @member {number} durationSec
         * @memberof lucar.Trip
         * @instance
         */
        Trip.prototype.durationSec = 0;

        /**
         * Trip endPos.
         * @member {lucar.ILocation|null|undefined} endPos
         * @memberof lucar.Trip
         * @instance
         */
        Trip.prototype.endPos = null;

        /**
         * Trip feeCent.
         * @member {number} feeCent
         * @memberof lucar.Trip
         * @instance
         */
        Trip.prototype.feeCent = 0;

        /**
         * Trip status.
         * @member {lucar.TripStatus} status
         * @memberof lucar.Trip
         * @instance
         */
        Trip.prototype.status = 0;

        /**
         * Creates a Trip message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof lucar.Trip
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {lucar.Trip} Trip
         */
        Trip.fromObject = function fromObject(object) {
            if (object instanceof $root.lucar.Trip)
                return object;
            let message = new $root.lucar.Trip();
            if (object.start != null)
                message.start = String(object.start);
            if (object.end != null)
                message.end = String(object.end);
            if (object.startPos != null) {
                if (typeof object.startPos !== "object")
                    throw TypeError(".lucar.Trip.startPos: object expected");
                message.startPos = $root.lucar.Location.fromObject(object.startPos);
            }
            if (object.pathLocations) {
                if (!Array.isArray(object.pathLocations))
                    throw TypeError(".lucar.Trip.pathLocations: array expected");
                message.pathLocations = [];
                for (let i = 0; i < object.pathLocations.length; ++i) {
                    if (typeof object.pathLocations[i] !== "object")
                        throw TypeError(".lucar.Trip.pathLocations: object expected");
                    message.pathLocations[i] = $root.lucar.Location.fromObject(object.pathLocations[i]);
                }
            }
            if (object.durationSec != null)
                message.durationSec = object.durationSec | 0;
            if (object.endPos != null) {
                if (typeof object.endPos !== "object")
                    throw TypeError(".lucar.Trip.endPos: object expected");
                message.endPos = $root.lucar.Location.fromObject(object.endPos);
            }
            if (object.feeCent != null)
                message.feeCent = object.feeCent | 0;
            switch (object.status) {
            default:
                if (typeof object.status === "number") {
                    message.status = object.status;
                    break;
                }
                break;
            case "TS_NOT_SPECIFIED":
            case 0:
                message.status = 0;
                break;
            case "NOT_STARTED":
            case 1:
                message.status = 1;
                break;
            case "IN_PROGRESS":
            case 2:
                message.status = 2;
                break;
            case "FINISHED":
            case 3:
                message.status = 3;
                break;
            case "PAID":
            case 4:
                message.status = 4;
                break;
            }
            return message;
        };

        /**
         * Creates a plain object from a Trip message. Also converts values to other types if specified.
         * @function toObject
         * @memberof lucar.Trip
         * @static
         * @param {lucar.Trip} message Trip
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Trip.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.arrays || options.defaults)
                object.pathLocations = [];
            if (options.defaults) {
                object.start = "";
                object.end = "";
                object.durationSec = 0;
                object.feeCent = 0;
                object.startPos = null;
                object.endPos = null;
                object.status = options.enums === String ? "TS_NOT_SPECIFIED" : 0;
            }
            if (message.start != null && message.hasOwnProperty("start"))
                object.start = message.start;
            if (message.end != null && message.hasOwnProperty("end"))
                object.end = message.end;
            if (message.durationSec != null && message.hasOwnProperty("durationSec"))
                object.durationSec = message.durationSec;
            if (message.feeCent != null && message.hasOwnProperty("feeCent"))
                object.feeCent = message.feeCent;
            if (message.startPos != null && message.hasOwnProperty("startPos"))
                object.startPos = $root.lucar.Location.toObject(message.startPos, options);
            if (message.endPos != null && message.hasOwnProperty("endPos"))
                object.endPos = $root.lucar.Location.toObject(message.endPos, options);
            if (message.pathLocations && message.pathLocations.length) {
                object.pathLocations = [];
                for (let j = 0; j < message.pathLocations.length; ++j)
                    object.pathLocations[j] = $root.lucar.Location.toObject(message.pathLocations[j], options);
            }
            if (message.status != null && message.hasOwnProperty("status"))
                object.status = options.enums === String ? $root.lucar.TripStatus[message.status] === undefined ? message.status : $root.lucar.TripStatus[message.status] : message.status;
            return object;
        };

        /**
         * Converts this Trip to JSON.
         * @function toJSON
         * @memberof lucar.Trip
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        Trip.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for Trip
         * @function getTypeUrl
         * @memberof lucar.Trip
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        Trip.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
            if (typeUrlPrefix === undefined) {
                typeUrlPrefix = "type.googleapis.com";
            }
            return typeUrlPrefix + "/lucar.Trip";
        };

        return Trip;
    })();

    lucar.GetTripRequest = (function() {

        /**
         * Properties of a GetTripRequest.
         * @memberof lucar
         * @interface IGetTripRequest
         * @property {string|null} [id] GetTripRequest id
         */

        /**
         * Constructs a new GetTripRequest.
         * @memberof lucar
         * @classdesc Represents a GetTripRequest.
         * @implements IGetTripRequest
         * @constructor
         * @param {lucar.IGetTripRequest=} [properties] Properties to set
         */
        function GetTripRequest(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * GetTripRequest id.
         * @member {string} id
         * @memberof lucar.GetTripRequest
         * @instance
         */
        GetTripRequest.prototype.id = "";

        /**
         * Creates a GetTripRequest message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof lucar.GetTripRequest
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {lucar.GetTripRequest} GetTripRequest
         */
        GetTripRequest.fromObject = function fromObject(object) {
            if (object instanceof $root.lucar.GetTripRequest)
                return object;
            let message = new $root.lucar.GetTripRequest();
            if (object.id != null)
                message.id = String(object.id);
            return message;
        };

        /**
         * Creates a plain object from a GetTripRequest message. Also converts values to other types if specified.
         * @function toObject
         * @memberof lucar.GetTripRequest
         * @static
         * @param {lucar.GetTripRequest} message GetTripRequest
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        GetTripRequest.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults)
                object.id = "";
            if (message.id != null && message.hasOwnProperty("id"))
                object.id = message.id;
            return object;
        };

        /**
         * Converts this GetTripRequest to JSON.
         * @function toJSON
         * @memberof lucar.GetTripRequest
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        GetTripRequest.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for GetTripRequest
         * @function getTypeUrl
         * @memberof lucar.GetTripRequest
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        GetTripRequest.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
            if (typeUrlPrefix === undefined) {
                typeUrlPrefix = "type.googleapis.com";
            }
            return typeUrlPrefix + "/lucar.GetTripRequest";
        };

        return GetTripRequest;
    })();

    lucar.GetTripResponse = (function() {

        /**
         * Properties of a GetTripResponse.
         * @memberof lucar
         * @interface IGetTripResponse
         * @property {string|null} [id] GetTripResponse id
         * @property {lucar.ITrip|null} [trip] GetTripResponse trip
         */

        /**
         * Constructs a new GetTripResponse.
         * @memberof lucar
         * @classdesc Represents a GetTripResponse.
         * @implements IGetTripResponse
         * @constructor
         * @param {lucar.IGetTripResponse=} [properties] Properties to set
         */
        function GetTripResponse(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * GetTripResponse id.
         * @member {string} id
         * @memberof lucar.GetTripResponse
         * @instance
         */
        GetTripResponse.prototype.id = "";

        /**
         * GetTripResponse trip.
         * @member {lucar.ITrip|null|undefined} trip
         * @memberof lucar.GetTripResponse
         * @instance
         */
        GetTripResponse.prototype.trip = null;

        /**
         * Creates a GetTripResponse message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof lucar.GetTripResponse
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {lucar.GetTripResponse} GetTripResponse
         */
        GetTripResponse.fromObject = function fromObject(object) {
            if (object instanceof $root.lucar.GetTripResponse)
                return object;
            let message = new $root.lucar.GetTripResponse();
            if (object.id != null)
                message.id = String(object.id);
            if (object.trip != null) {
                if (typeof object.trip !== "object")
                    throw TypeError(".lucar.GetTripResponse.trip: object expected");
                message.trip = $root.lucar.Trip.fromObject(object.trip);
            }
            return message;
        };

        /**
         * Creates a plain object from a GetTripResponse message. Also converts values to other types if specified.
         * @function toObject
         * @memberof lucar.GetTripResponse
         * @static
         * @param {lucar.GetTripResponse} message GetTripResponse
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        GetTripResponse.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults) {
                object.id = "";
                object.trip = null;
            }
            if (message.id != null && message.hasOwnProperty("id"))
                object.id = message.id;
            if (message.trip != null && message.hasOwnProperty("trip"))
                object.trip = $root.lucar.Trip.toObject(message.trip, options);
            return object;
        };

        /**
         * Converts this GetTripResponse to JSON.
         * @function toJSON
         * @memberof lucar.GetTripResponse
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        GetTripResponse.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for GetTripResponse
         * @function getTypeUrl
         * @memberof lucar.GetTripResponse
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        GetTripResponse.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
            if (typeUrlPrefix === undefined) {
                typeUrlPrefix = "type.googleapis.com";
            }
            return typeUrlPrefix + "/lucar.GetTripResponse";
        };

        return GetTripResponse;
    })();

    lucar.TripService = (function() {

        /**
         * Constructs a new TripService service.
         * @memberof lucar
         * @classdesc Represents a TripService
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function TripService(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (TripService.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = TripService;

        /**
         * Callback as used by {@link lucar.TripService#getTrip}.
         * @memberof lucar.TripService
         * @typedef GetTripCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {lucar.GetTripResponse} [response] GetTripResponse
         */

        /**
         * Calls GetTrip.
         * @function getTrip
         * @memberof lucar.TripService
         * @instance
         * @param {lucar.IGetTripRequest} request GetTripRequest message or plain object
         * @param {lucar.TripService.GetTripCallback} callback Node-style callback called with the error, if any, and GetTripResponse
         * @returns {undefined}
         * @variation 1
         */
        Object.defineProperty(TripService.prototype.getTrip = function getTrip(request, callback) {
            return this.rpcCall(getTrip, $root.lucar.GetTripRequest, $root.lucar.GetTripResponse, request, callback);
        }, "name", { value: "GetTrip" });

        /**
         * Calls GetTrip.
         * @function getTrip
         * @memberof lucar.TripService
         * @instance
         * @param {lucar.IGetTripRequest} request GetTripRequest message or plain object
         * @returns {Promise<lucar.GetTripResponse>} Promise
         * @variation 2
         */

        return TripService;
    })();

    return lucar;
})();