// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: soccer.proto

package soccer

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SoccerSimulatorClient is the client API for SoccerSimulator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SoccerSimulatorClient interface {
	PlaceBet(ctx context.Context, in *Bet, opts ...grpc.CallOption) (*BetUpdateResponse, error)
	GetMatch(ctx context.Context, in *GetMatchRequest, opts ...grpc.CallOption) (*Match, error)
	UpdateTeam(ctx context.Context, in *Team, opts ...grpc.CallOption) (*AddTeamResponse, error)
	AddTeam(ctx context.Context, in *AddTeamRequest, opts ...grpc.CallOption) (*AddTeamResponse, error)
	GetBets(ctx context.Context, in *GetBetsRequest, opts ...grpc.CallOption) (*GetBetsResponse, error)
	DeleteBet(ctx context.Context, in *DeleteBetRequest, opts ...grpc.CallOption) (*BetUpdateResponse, error)
	AddLeague(ctx context.Context, in *AddLeagueRequest, opts ...grpc.CallOption) (*AddLeagueResponse, error)
	UpdateBet(ctx context.Context, in *UpdateBetRequest, opts ...grpc.CallOption) (*BetUpdateResponse, error)
	GetTeams(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetTeamsResponse, error)
	GetMatches(ctx context.Context, in *GetMatchesRequest, opts ...grpc.CallOption) (*GetMatchesResponse, error)
	DeleteTeam(ctx context.Context, in *DeleteTeamRequest, opts ...grpc.CallOption) (*DeleteTeamResponse, error)
	AddOutcomes(ctx context.Context, in *AddOutcomesRequest, opts ...grpc.CallOption) (*AddOutcomesResponse, error)
	GetLeagues(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetLeaguesResponse, error)
	DeleteLeague(ctx context.Context, in *DeleteLeagueRequest, opts ...grpc.CallOption) (*DeleteLeagueResponse, error)
	SubscribeSoccer(ctx context.Context, in *SubscribeSoccerRequest, opts ...grpc.CallOption) (*SubscribeSoccerResponse, error)
	GetSoccerSettings(ctx context.Context, in *GetSoccerSettingsRequest, opts ...grpc.CallOption) (*SoccerSimulatorAdmin, error)
	UpdateSoccerSettings(ctx context.Context, in *SoccerSimulatorAdmin, opts ...grpc.CallOption) (*UpdateSoccerSettingsResponse, error)
}

type soccerSimulatorClient struct {
	cc grpc.ClientConnInterface
}

func NewSoccerSimulatorClient(cc grpc.ClientConnInterface) SoccerSimulatorClient {
	return &soccerSimulatorClient{cc}
}

func (c *soccerSimulatorClient) PlaceBet(ctx context.Context, in *Bet, opts ...grpc.CallOption) (*BetUpdateResponse, error) {
	out := new(BetUpdateResponse)
	err := c.cc.Invoke(ctx, "/SoccerSimulator/PlaceBet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *soccerSimulatorClient) GetMatch(ctx context.Context, in *GetMatchRequest, opts ...grpc.CallOption) (*Match, error) {
	out := new(Match)
	err := c.cc.Invoke(ctx, "/SoccerSimulator/GetMatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *soccerSimulatorClient) UpdateTeam(ctx context.Context, in *Team, opts ...grpc.CallOption) (*AddTeamResponse, error) {
	out := new(AddTeamResponse)
	err := c.cc.Invoke(ctx, "/SoccerSimulator/UpdateTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *soccerSimulatorClient) AddTeam(ctx context.Context, in *AddTeamRequest, opts ...grpc.CallOption) (*AddTeamResponse, error) {
	out := new(AddTeamResponse)
	err := c.cc.Invoke(ctx, "/SoccerSimulator/AddTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *soccerSimulatorClient) GetBets(ctx context.Context, in *GetBetsRequest, opts ...grpc.CallOption) (*GetBetsResponse, error) {
	out := new(GetBetsResponse)
	err := c.cc.Invoke(ctx, "/SoccerSimulator/GetBets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *soccerSimulatorClient) DeleteBet(ctx context.Context, in *DeleteBetRequest, opts ...grpc.CallOption) (*BetUpdateResponse, error) {
	out := new(BetUpdateResponse)
	err := c.cc.Invoke(ctx, "/SoccerSimulator/DeleteBet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *soccerSimulatorClient) AddLeague(ctx context.Context, in *AddLeagueRequest, opts ...grpc.CallOption) (*AddLeagueResponse, error) {
	out := new(AddLeagueResponse)
	err := c.cc.Invoke(ctx, "/SoccerSimulator/AddLeague", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *soccerSimulatorClient) UpdateBet(ctx context.Context, in *UpdateBetRequest, opts ...grpc.CallOption) (*BetUpdateResponse, error) {
	out := new(BetUpdateResponse)
	err := c.cc.Invoke(ctx, "/SoccerSimulator/UpdateBet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *soccerSimulatorClient) GetTeams(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetTeamsResponse, error) {
	out := new(GetTeamsResponse)
	err := c.cc.Invoke(ctx, "/SoccerSimulator/GetTeams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *soccerSimulatorClient) GetMatches(ctx context.Context, in *GetMatchesRequest, opts ...grpc.CallOption) (*GetMatchesResponse, error) {
	out := new(GetMatchesResponse)
	err := c.cc.Invoke(ctx, "/SoccerSimulator/GetMatches", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *soccerSimulatorClient) DeleteTeam(ctx context.Context, in *DeleteTeamRequest, opts ...grpc.CallOption) (*DeleteTeamResponse, error) {
	out := new(DeleteTeamResponse)
	err := c.cc.Invoke(ctx, "/SoccerSimulator/DeleteTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *soccerSimulatorClient) AddOutcomes(ctx context.Context, in *AddOutcomesRequest, opts ...grpc.CallOption) (*AddOutcomesResponse, error) {
	out := new(AddOutcomesResponse)
	err := c.cc.Invoke(ctx, "/SoccerSimulator/AddOutcomes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *soccerSimulatorClient) GetLeagues(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetLeaguesResponse, error) {
	out := new(GetLeaguesResponse)
	err := c.cc.Invoke(ctx, "/SoccerSimulator/GetLeagues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *soccerSimulatorClient) DeleteLeague(ctx context.Context, in *DeleteLeagueRequest, opts ...grpc.CallOption) (*DeleteLeagueResponse, error) {
	out := new(DeleteLeagueResponse)
	err := c.cc.Invoke(ctx, "/SoccerSimulator/DeleteLeague", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *soccerSimulatorClient) SubscribeSoccer(ctx context.Context, in *SubscribeSoccerRequest, opts ...grpc.CallOption) (*SubscribeSoccerResponse, error) {
	out := new(SubscribeSoccerResponse)
	err := c.cc.Invoke(ctx, "/SoccerSimulator/SubscribeSoccer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *soccerSimulatorClient) GetSoccerSettings(ctx context.Context, in *GetSoccerSettingsRequest, opts ...grpc.CallOption) (*SoccerSimulatorAdmin, error) {
	out := new(SoccerSimulatorAdmin)
	err := c.cc.Invoke(ctx, "/SoccerSimulator/GetSoccerSettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *soccerSimulatorClient) UpdateSoccerSettings(ctx context.Context, in *SoccerSimulatorAdmin, opts ...grpc.CallOption) (*UpdateSoccerSettingsResponse, error) {
	out := new(UpdateSoccerSettingsResponse)
	err := c.cc.Invoke(ctx, "/SoccerSimulator/UpdateSoccerSettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SoccerSimulatorServer is the server API for SoccerSimulator service.
// All implementations should embed UnimplementedSoccerSimulatorServer
// for forward compatibility
type SoccerSimulatorServer interface {
	PlaceBet(context.Context, *Bet) (*BetUpdateResponse, error)
	GetMatch(context.Context, *GetMatchRequest) (*Match, error)
	UpdateTeam(context.Context, *Team) (*AddTeamResponse, error)
	AddTeam(context.Context, *AddTeamRequest) (*AddTeamResponse, error)
	GetBets(context.Context, *GetBetsRequest) (*GetBetsResponse, error)
	DeleteBet(context.Context, *DeleteBetRequest) (*BetUpdateResponse, error)
	AddLeague(context.Context, *AddLeagueRequest) (*AddLeagueResponse, error)
	UpdateBet(context.Context, *UpdateBetRequest) (*BetUpdateResponse, error)
	GetTeams(context.Context, *empty.Empty) (*GetTeamsResponse, error)
	GetMatches(context.Context, *GetMatchesRequest) (*GetMatchesResponse, error)
	DeleteTeam(context.Context, *DeleteTeamRequest) (*DeleteTeamResponse, error)
	AddOutcomes(context.Context, *AddOutcomesRequest) (*AddOutcomesResponse, error)
	GetLeagues(context.Context, *empty.Empty) (*GetLeaguesResponse, error)
	DeleteLeague(context.Context, *DeleteLeagueRequest) (*DeleteLeagueResponse, error)
	SubscribeSoccer(context.Context, *SubscribeSoccerRequest) (*SubscribeSoccerResponse, error)
	GetSoccerSettings(context.Context, *GetSoccerSettingsRequest) (*SoccerSimulatorAdmin, error)
	UpdateSoccerSettings(context.Context, *SoccerSimulatorAdmin) (*UpdateSoccerSettingsResponse, error)
}

// UnimplementedSoccerSimulatorServer should be embedded to have forward compatible implementations.
type UnimplementedSoccerSimulatorServer struct {
}

func (UnimplementedSoccerSimulatorServer) PlaceBet(context.Context, *Bet) (*BetUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceBet not implemented")
}
func (UnimplementedSoccerSimulatorServer) GetMatch(context.Context, *GetMatchRequest) (*Match, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMatch not implemented")
}
func (UnimplementedSoccerSimulatorServer) UpdateTeam(context.Context, *Team) (*AddTeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTeam not implemented")
}
func (UnimplementedSoccerSimulatorServer) AddTeam(context.Context, *AddTeamRequest) (*AddTeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTeam not implemented")
}
func (UnimplementedSoccerSimulatorServer) GetBets(context.Context, *GetBetsRequest) (*GetBetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBets not implemented")
}
func (UnimplementedSoccerSimulatorServer) DeleteBet(context.Context, *DeleteBetRequest) (*BetUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBet not implemented")
}
func (UnimplementedSoccerSimulatorServer) AddLeague(context.Context, *AddLeagueRequest) (*AddLeagueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddLeague not implemented")
}
func (UnimplementedSoccerSimulatorServer) UpdateBet(context.Context, *UpdateBetRequest) (*BetUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBet not implemented")
}
func (UnimplementedSoccerSimulatorServer) GetTeams(context.Context, *empty.Empty) (*GetTeamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeams not implemented")
}
func (UnimplementedSoccerSimulatorServer) GetMatches(context.Context, *GetMatchesRequest) (*GetMatchesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMatches not implemented")
}
func (UnimplementedSoccerSimulatorServer) DeleteTeam(context.Context, *DeleteTeamRequest) (*DeleteTeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTeam not implemented")
}
func (UnimplementedSoccerSimulatorServer) AddOutcomes(context.Context, *AddOutcomesRequest) (*AddOutcomesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOutcomes not implemented")
}
func (UnimplementedSoccerSimulatorServer) GetLeagues(context.Context, *empty.Empty) (*GetLeaguesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLeagues not implemented")
}
func (UnimplementedSoccerSimulatorServer) DeleteLeague(context.Context, *DeleteLeagueRequest) (*DeleteLeagueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLeague not implemented")
}
func (UnimplementedSoccerSimulatorServer) SubscribeSoccer(context.Context, *SubscribeSoccerRequest) (*SubscribeSoccerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubscribeSoccer not implemented")
}
func (UnimplementedSoccerSimulatorServer) GetSoccerSettings(context.Context, *GetSoccerSettingsRequest) (*SoccerSimulatorAdmin, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSoccerSettings not implemented")
}
func (UnimplementedSoccerSimulatorServer) UpdateSoccerSettings(context.Context, *SoccerSimulatorAdmin) (*UpdateSoccerSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSoccerSettings not implemented")
}

// UnsafeSoccerSimulatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SoccerSimulatorServer will
// result in compilation errors.
type UnsafeSoccerSimulatorServer interface {
	mustEmbedUnimplementedSoccerSimulatorServer()
}

func RegisterSoccerSimulatorServer(s grpc.ServiceRegistrar, srv SoccerSimulatorServer) {
	s.RegisterService(&SoccerSimulator_ServiceDesc, srv)
}

func _SoccerSimulator_PlaceBet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Bet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SoccerSimulatorServer).PlaceBet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SoccerSimulator/PlaceBet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SoccerSimulatorServer).PlaceBet(ctx, req.(*Bet))
	}
	return interceptor(ctx, in, info, handler)
}

func _SoccerSimulator_GetMatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SoccerSimulatorServer).GetMatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SoccerSimulator/GetMatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SoccerSimulatorServer).GetMatch(ctx, req.(*GetMatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SoccerSimulator_UpdateTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Team)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SoccerSimulatorServer).UpdateTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SoccerSimulator/UpdateTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SoccerSimulatorServer).UpdateTeam(ctx, req.(*Team))
	}
	return interceptor(ctx, in, info, handler)
}

func _SoccerSimulator_AddTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SoccerSimulatorServer).AddTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SoccerSimulator/AddTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SoccerSimulatorServer).AddTeam(ctx, req.(*AddTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SoccerSimulator_GetBets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SoccerSimulatorServer).GetBets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SoccerSimulator/GetBets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SoccerSimulatorServer).GetBets(ctx, req.(*GetBetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SoccerSimulator_DeleteBet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SoccerSimulatorServer).DeleteBet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SoccerSimulator/DeleteBet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SoccerSimulatorServer).DeleteBet(ctx, req.(*DeleteBetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SoccerSimulator_AddLeague_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddLeagueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SoccerSimulatorServer).AddLeague(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SoccerSimulator/AddLeague",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SoccerSimulatorServer).AddLeague(ctx, req.(*AddLeagueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SoccerSimulator_UpdateBet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SoccerSimulatorServer).UpdateBet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SoccerSimulator/UpdateBet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SoccerSimulatorServer).UpdateBet(ctx, req.(*UpdateBetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SoccerSimulator_GetTeams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SoccerSimulatorServer).GetTeams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SoccerSimulator/GetTeams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SoccerSimulatorServer).GetTeams(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SoccerSimulator_GetMatches_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMatchesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SoccerSimulatorServer).GetMatches(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SoccerSimulator/GetMatches",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SoccerSimulatorServer).GetMatches(ctx, req.(*GetMatchesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SoccerSimulator_DeleteTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SoccerSimulatorServer).DeleteTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SoccerSimulator/DeleteTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SoccerSimulatorServer).DeleteTeam(ctx, req.(*DeleteTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SoccerSimulator_AddOutcomes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddOutcomesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SoccerSimulatorServer).AddOutcomes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SoccerSimulator/AddOutcomes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SoccerSimulatorServer).AddOutcomes(ctx, req.(*AddOutcomesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SoccerSimulator_GetLeagues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SoccerSimulatorServer).GetLeagues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SoccerSimulator/GetLeagues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SoccerSimulatorServer).GetLeagues(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SoccerSimulator_DeleteLeague_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLeagueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SoccerSimulatorServer).DeleteLeague(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SoccerSimulator/DeleteLeague",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SoccerSimulatorServer).DeleteLeague(ctx, req.(*DeleteLeagueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SoccerSimulator_SubscribeSoccer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeSoccerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SoccerSimulatorServer).SubscribeSoccer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SoccerSimulator/SubscribeSoccer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SoccerSimulatorServer).SubscribeSoccer(ctx, req.(*SubscribeSoccerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SoccerSimulator_GetSoccerSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSoccerSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SoccerSimulatorServer).GetSoccerSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SoccerSimulator/GetSoccerSettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SoccerSimulatorServer).GetSoccerSettings(ctx, req.(*GetSoccerSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SoccerSimulator_UpdateSoccerSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SoccerSimulatorAdmin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SoccerSimulatorServer).UpdateSoccerSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SoccerSimulator/UpdateSoccerSettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SoccerSimulatorServer).UpdateSoccerSettings(ctx, req.(*SoccerSimulatorAdmin))
	}
	return interceptor(ctx, in, info, handler)
}

// SoccerSimulator_ServiceDesc is the grpc.ServiceDesc for SoccerSimulator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SoccerSimulator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SoccerSimulator",
	HandlerType: (*SoccerSimulatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PlaceBet",
			Handler:    _SoccerSimulator_PlaceBet_Handler,
		},
		{
			MethodName: "GetMatch",
			Handler:    _SoccerSimulator_GetMatch_Handler,
		},
		{
			MethodName: "UpdateTeam",
			Handler:    _SoccerSimulator_UpdateTeam_Handler,
		},
		{
			MethodName: "AddTeam",
			Handler:    _SoccerSimulator_AddTeam_Handler,
		},
		{
			MethodName: "GetBets",
			Handler:    _SoccerSimulator_GetBets_Handler,
		},
		{
			MethodName: "DeleteBet",
			Handler:    _SoccerSimulator_DeleteBet_Handler,
		},
		{
			MethodName: "AddLeague",
			Handler:    _SoccerSimulator_AddLeague_Handler,
		},
		{
			MethodName: "UpdateBet",
			Handler:    _SoccerSimulator_UpdateBet_Handler,
		},
		{
			MethodName: "GetTeams",
			Handler:    _SoccerSimulator_GetTeams_Handler,
		},
		{
			MethodName: "GetMatches",
			Handler:    _SoccerSimulator_GetMatches_Handler,
		},
		{
			MethodName: "DeleteTeam",
			Handler:    _SoccerSimulator_DeleteTeam_Handler,
		},
		{
			MethodName: "AddOutcomes",
			Handler:    _SoccerSimulator_AddOutcomes_Handler,
		},
		{
			MethodName: "GetLeagues",
			Handler:    _SoccerSimulator_GetLeagues_Handler,
		},
		{
			MethodName: "DeleteLeague",
			Handler:    _SoccerSimulator_DeleteLeague_Handler,
		},
		{
			MethodName: "SubscribeSoccer",
			Handler:    _SoccerSimulator_SubscribeSoccer_Handler,
		},
		{
			MethodName: "GetSoccerSettings",
			Handler:    _SoccerSimulator_GetSoccerSettings_Handler,
		},
		{
			MethodName: "UpdateSoccerSettings",
			Handler:    _SoccerSimulator_UpdateSoccerSettings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "soccer.proto",
}
