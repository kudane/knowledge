abstract class WithInputWithOutput<TInput, TOutput> {
    abstract executeAsync(input: TInput): Promise<TOutput>;
}

abstract class WithInputNoOutput<TInput> {
    abstract executeAsync(input: TInput): Promise<void>;
}

abstract class NoInputWithOutput<TOutput> {
    abstract executeAsync(): Promise<TOutput>;
}

abstract class NoInputNoOutput {
    abstract executeAsync(): Promise<void>;
}

export abstract class ServiceHandlerAsync {
    static WithInput = class {
        static WithOutput = WithInputWithOutput;
        static NoOutput = WithInputNoOutput;
    };

    static NoInput = class {
        static WithOutput = NoInputWithOutput;
        static NoOutput = NoInputNoOutput;
    };
}