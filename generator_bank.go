package chinaid

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"

	"github.com/mritd/chinaid/metadata"
)

// BankNo 返回随机银行卡号，银行卡号符合LUHN 算法并且有正确的卡 bin 前缀
func BankNo() string {
	// 随机选中银行卡卡头
	bank := metadata.CardBins[rand.Intn(len(metadata.CardBins))]
	// 获取 卡前缀(cardBin)
	prefixes := bank.Prefixes
	// 获取当前银行卡正确长度
	cardNoLength := bank.Length
	// 生成 长度-1 位卡号
	preCardNo := strconv.Itoa(prefixes[randInt(0, len(prefixes))]) + fmt.Sprintf("%0*d", cardNoLength-7, randInt64(0, int64(math.Pow10(cardNoLength-7))))
	// LUHN 算法处理
	return LUHNProcess(preCardNo)
}

// LUHNProcess 通过 LUHN 合成卡号处理给定的银行卡号
func LUHNProcess(preCardNo string) string {
	checkSum := 0
	tmpCardNo := reverseString(preCardNo)
	for i, s := range tmpCardNo {
		// 数据层确保卡号正确
		tmp, _ := strconv.Atoi(string(s))
		// 由于卡号实际少了一位，所以反转后卡号第一位一定为偶数位
		// 同时 i 正好也是偶数，此时 i 将和卡号奇偶位同步
		if i%2 == 0 {
			// 偶数位 *2 是否为两位数(>9)
			if tmp*2 > 9 {
				// 如果为两位数则 -9
				checkSum += tmp*2 - 9
			} else {
				// 否则直接相加即可
				checkSum += tmp * 2
			}
		} else {
			// 奇数位直接相加
			checkSum += tmp
		}
	}
	if checkSum%10 != 0 {
		return preCardNo + strconv.Itoa(10-checkSum%10)
	} else {
		// 如果不巧生成的前 卡长度-1 位正好符合 LUHN 算法
		// 那么需要递归重新生成(需要符合 cardBind 中卡号长度)
		return BankNo()
	}
}
