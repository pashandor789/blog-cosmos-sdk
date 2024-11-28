package module

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	blogv1 "github.com/blog/api/v1"
)

func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: blogv1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "GetPost",
					Use:       "get-post index",
					Short:     "Get the current value of the post at index",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "index"},
					},
				},
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service: blogv1.Msg_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "CreatePost",
					Use:       "create index black red",
					Short:     "Creates a new blog. post at the index for the black and red players",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "index"},
						{ProtoField: "text"},
					},
				},
			},
		},
	}
}
