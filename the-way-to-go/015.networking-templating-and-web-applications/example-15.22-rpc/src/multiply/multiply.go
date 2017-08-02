package multiply

type Args struct {
    N, M int
}

func (a *Args) Multiply(args *Args, reply *int) error {
    *reply = args.N * args.M
    return  nil
}
