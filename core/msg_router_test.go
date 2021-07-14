package core

import (
	"encoding/hex"
	"testing"
)

func TestGetSignBytes(t *testing.T) {
	//addr := "t3sxzvhng4nqko7nv26hgkl53tak5kzk6pjgmrevnehk6cidrafhtcu53n6hpyj6hr4m2c362hxrtpacndicpa"
	mtblockData := "904300e8078158609117ac930d5e257fd6117df27a552e65ec60ea759133cf83e784bcaf936ba241593f1e24553c027d" +
		"69b6b4dd57a99df2030c2154f8bc3bb3600f3c5de6f9eda618e31a20f853af183007cc4beff10bbaf6eba21fa18dfe0610a1001e6c" +
		"14bb4d82045860b662e9f11b3cd549d12b4682bb108b3b18a00756f42887cf269b57ab14b4b209d3f1c0d09cb88ec38d438f4554a2" +
		"27e31369c941b8aabe4fed7fc6e74aaa59fff9ac72863919667341db41b50798c62e81df9518c7a0ca8a10a1b3737e353458808182" +
		"0058c0b71215ac347e816e200e8d06c7d173f3b0ab44e4834385520920ba23dc5f9191de38f4e47b9ecd79ce8d1d22acecbc0ab61b" +
		"944a83d819d15c7b4a82d763e8f0fe662b53b2a8437aaade355ee918049f4cd04b0899512c09820eb5e5f1b3bc3e03ce5effac39e4" +
		"339ae7d9390cc332c99f78329d9794d0b27721ad29bd3aab249893ecc1882c1ebe41b7b8854ccff515b9cdc6590cb6dece0d77215f" +
		"a57bdb718fe53da0f5509d70bfa140f2d56750f5f244e17ff7cec611999e7f61c73f6d2e81d82a5827000171a0e402207eb6789dd0" +
		"560a9b323ade4ffd545a64cd3f364396d47ec766b43044d1ea0ef74400070380184fd82a5827000171a0e40220729151866016f09d" +
		"350cc11a3a05074fab990628cdcdf44c8dd3a6f702a30ee8d82a5827000171a0e40220e5658b3d18cd06e1db9015b4b0ec55c123a2" +
		"4d5be1ea24d83938c5b8397b4f2fd82a5827000171a0e4022098307faeabdd208c1321be4c1ec84fbb952cd8424c24db0d5a96c803" +
		"2fb2ce10586102c0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000" +
		"00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001a6034975cf60043000a3a"
	mtblockBytes, err := hex.DecodeString(mtblockData)
	if err != nil {
		t.Fatalf("block meta decode err:%s", err)
	}
	_, _, err = GetSignBytes(mtblockBytes, MsgMeta{Type: MTBlock})
	if err != nil {
		t.Fatal(err)
	}
}
