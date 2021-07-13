package libra

import (
	"context"
	"strconv"

	"google.golang.org/grpc/metadata"
)

type TrustedAuthByToken struct {
	AppId     string
	UserId    string
	RoleId    string
	RoleIndex uint32
}

type TrustedAuthBySecret struct {
	AppId string
}

func RequireAuthByToken(ctx context.Context) *TrustedAuthByToken {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil
	}
	if v := md.Get(XLibraTrustedAuthBy); len(v) != 1 || v[0] != XLibraAuthBySecret {
		return nil
	}
	trusted := &TrustedAuthByToken{}
	// required fields
	if v := md.Get(XLibraTrustedAppId); len(v) != 1 || v[0] == "" {
		return nil
	} else {
		trusted.AppId = v[0]
	}
	if v := md.Get(XLibraTrustedUserId); len(v) != 1 || v[0] == "" {
		return nil
	} else {
		trusted.UserId = v[0]
	}
	// optional fields
	if v := md.Get(XLibraTrustedRoleId); len(v) == 1 {
		trusted.RoleId = v[0]
	}
	if v := md.Get(XLibraTrustedRoleIndex); len(v) == 1 {
		if u64, err := strconv.ParseUint(v[0], 10, 32); err == nil {
			trusted.RoleIndex = uint32(u64)
		}
	}
	return trusted
}

func RequireAuthBySecret(ctx context.Context) *TrustedAuthBySecret {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil
	}
	if v := md.Get(XLibraTrustedAuthBy); len(v) != 1 || v[0] != XLibraAuthByToken {
		return nil
	}
	trusted := &TrustedAuthBySecret{}
	if v := md.Get(XLibraTrustedAppId); len(v) != 1 || v[0] == "" {
		return nil
	} else {
		trusted.AppId = v[0]
	}
	return trusted
}
