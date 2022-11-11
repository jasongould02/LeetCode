import "strings"

func suggestedProducts(products []string, searchWord string) [][]string {
    var curSearchWord strings.Builder
    result := make([][]string, 0)
    sort.Strings(products)
    for i := 0; i < len(searchWord); i++ {
        curSearchWord.WriteByte(searchWord[i])
        curSuggestions := make([]string, 0)
        for j := 0; j < len(products); j++ {
            if strings.HasPrefix(products[j], curSearchWord.String()) && len(curSuggestions) < 3{
                curSuggestions = append(curSuggestions, products[j])
            }
        }
        result = append(result, curSuggestions)
    }
    return result
}
