import * as $protobuf from "protobufjs";
import Long = require("long");
/** Namespace lucar. */
export namespace lucar {

    /** Properties of a Location. */
    interface ILocation {

        /** Location latitude */
        latitude?: (number|null);

        /** Location longitude */
        longitude?: (number|null);
    }

    /** Represents a Location. */
    class Location implements ILocation {

        /**
         * Constructs a new Location.
         * @param [properties] Properties to set
         */
        constructor(properties?: lucar.ILocation);

        /** Location latitude. */
        public latitude: number;

        /** Location longitude. */
        public longitude: number;

        /**
         * Creates a Location message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns Location
         */
        public static fromObject(object: { [k: string]: any }): lucar.Location;

        /**
         * Creates a plain object from a Location message. Also converts values to other types if specified.
         * @param message Location
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: lucar.Location, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this Location to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };

        /**
         * Gets the default type url for Location
         * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns The default type url
         */
        public static getTypeUrl(typeUrlPrefix?: string): string;
    }

    /** TripStatus enum. */
    enum TripStatus {
        TS_NOT_SPECIFIED = 0,
        NOT_STARTED = 1,
        IN_PROGRESS = 2,
        FINISHED = 3,
        PAID = 4
    }

    /** Properties of a Trip. */
    interface ITrip {

        /** Trip start */
        start?: (string|null);

        /** Trip end */
        end?: (string|null);

        /** Trip startPos */
        startPos?: (lucar.ILocation|null);

        /** Trip pathLocations */
        pathLocations?: (lucar.ILocation[]|null);

        /** Trip durationSec */
        durationSec?: (number|null);

        /** Trip endPos */
        endPos?: (lucar.ILocation|null);

        /** Trip feeCent */
        feeCent?: (number|null);

        /** Trip status */
        status?: (lucar.TripStatus|null);
    }

    /** Represents a Trip. */
    class Trip implements ITrip {

        /**
         * Constructs a new Trip.
         * @param [properties] Properties to set
         */
        constructor(properties?: lucar.ITrip);

        /** Trip start. */
        public start: string;

        /** Trip end. */
        public end: string;

        /** Trip startPos. */
        public startPos?: (lucar.ILocation|null);

        /** Trip pathLocations. */
        public pathLocations: lucar.ILocation[];

        /** Trip durationSec. */
        public durationSec: number;

        /** Trip endPos. */
        public endPos?: (lucar.ILocation|null);

        /** Trip feeCent. */
        public feeCent: number;

        /** Trip status. */
        public status: lucar.TripStatus;

        /**
         * Creates a Trip message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns Trip
         */
        public static fromObject(object: { [k: string]: any }): lucar.Trip;

        /**
         * Creates a plain object from a Trip message. Also converts values to other types if specified.
         * @param message Trip
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: lucar.Trip, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this Trip to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };

        /**
         * Gets the default type url for Trip
         * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns The default type url
         */
        public static getTypeUrl(typeUrlPrefix?: string): string;
    }

    /** Properties of a GetTripRequest. */
    interface IGetTripRequest {

        /** GetTripRequest id */
        id?: (string|null);
    }

    /** Represents a GetTripRequest. */
    class GetTripRequest implements IGetTripRequest {

        /**
         * Constructs a new GetTripRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: lucar.IGetTripRequest);

        /** GetTripRequest id. */
        public id: string;

        /**
         * Creates a GetTripRequest message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns GetTripRequest
         */
        public static fromObject(object: { [k: string]: any }): lucar.GetTripRequest;

        /**
         * Creates a plain object from a GetTripRequest message. Also converts values to other types if specified.
         * @param message GetTripRequest
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: lucar.GetTripRequest, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this GetTripRequest to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };

        /**
         * Gets the default type url for GetTripRequest
         * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns The default type url
         */
        public static getTypeUrl(typeUrlPrefix?: string): string;
    }

    /** Properties of a GetTripResponse. */
    interface IGetTripResponse {

        /** GetTripResponse id */
        id?: (string|null);

        /** GetTripResponse trip */
        trip?: (lucar.ITrip|null);
    }

    /** Represents a GetTripResponse. */
    class GetTripResponse implements IGetTripResponse {

        /**
         * Constructs a new GetTripResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: lucar.IGetTripResponse);

        /** GetTripResponse id. */
        public id: string;

        /** GetTripResponse trip. */
        public trip?: (lucar.ITrip|null);

        /**
         * Creates a GetTripResponse message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns GetTripResponse
         */
        public static fromObject(object: { [k: string]: any }): lucar.GetTripResponse;

        /**
         * Creates a plain object from a GetTripResponse message. Also converts values to other types if specified.
         * @param message GetTripResponse
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: lucar.GetTripResponse, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this GetTripResponse to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };

        /**
         * Gets the default type url for GetTripResponse
         * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns The default type url
         */
        public static getTypeUrl(typeUrlPrefix?: string): string;
    }

    /** Represents a TripService */
    class TripService extends $protobuf.rpc.Service {

        /**
         * Constructs a new TripService service.
         * @param rpcImpl RPC implementation
         * @param [requestDelimited=false] Whether requests are length-delimited
         * @param [responseDelimited=false] Whether responses are length-delimited
         */
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);

        /**
         * Calls GetTrip.
         * @param request GetTripRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and GetTripResponse
         */
        public getTrip(request: lucar.IGetTripRequest, callback: lucar.TripService.GetTripCallback): void;

        /**
         * Calls GetTrip.
         * @param request GetTripRequest message or plain object
         * @returns Promise
         */
        public getTrip(request: lucar.IGetTripRequest): Promise<lucar.GetTripResponse>;
    }

    namespace TripService {

        /**
         * Callback as used by {@link lucar.TripService#getTrip}.
         * @param error Error, if any
         * @param [response] GetTripResponse
         */
        type GetTripCallback = (error: (Error|null), response?: lucar.GetTripResponse) => void;
    }
}
