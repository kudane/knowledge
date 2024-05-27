abstract class WithRequestWithResponse<TRequest, TResponse> {
    abstract handleAsync(request: TRequest): Promise<TResponse>;
}

abstract class WithRequestNoResponse<TRequest> {
    abstract handleAsync(request: TRequest): Promise<void>;
}

abstract class NoRequestWithResponse<TResponse> {
    abstract handleAsync(): Promise<TResponse>;
}

abstract class NoRequestNoResponse {
    abstract handleAsync(): Promise<void>;
}

export abstract class ApiHandlerAsync {
    static WithRequest = class {
        static WithResponse = WithRequestWithResponse;
        static NoResponse = WithRequestNoResponse;
    };

    static NoRequest = class {
        static WithResponse = NoRequestWithResponse;
        static NoResponse = NoRequestNoResponse;
    };
}