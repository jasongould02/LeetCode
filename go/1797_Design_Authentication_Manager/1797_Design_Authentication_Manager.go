type AuthenticationManager struct {
    timeToLive int
    m map[string]int
}


func Constructor(timeToLive int) AuthenticationManager {
    var am AuthenticationManager
    am.timeToLive = timeToLive
    am.m = make(map[string]int)
    return am
}


func (this *AuthenticationManager) Generate(tokenId string, currentTime int)  {
    this.m[tokenId] = currentTime + this.timeToLive
}


func (this *AuthenticationManager) Renew(tokenId string, currentTime int)  {
    _, exists := this.m[tokenId]
    if exists && this.m[tokenId] > currentTime {
        this.m[tokenId] = currentTime + this.timeToLive
    }
}


func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
    total := 0
    for _, v := range this.m {
        if v > currentTime {
            total += 1
        }
    }
    return total
}


/**
 * Your AuthenticationManager object will be instantiated and called as such:
 * obj := Constructor(timeToLive);
 * obj.Generate(tokenId,currentTime);
 * obj.Renew(tokenId,currentTime);
 * param_3 := obj.CountUnexpiredTokens(currentTime);
 */
