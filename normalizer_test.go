package main
import ("testing"; "encoding/json")
func TestNestedEvents(t *testing.T) {
    rows:=[]any{map[string]any{"id":"sample-a"}}
    value:=map[string]any{"data":map[string]any{"data":rows}}
    got:=findOccurrenceArray(value)
    if len(got)!=1 || valueString(objectValue(got[0],"id"))!="sample-a" {t.Fatal("wrong nested result")}
}
func TestEmptyAndUnknown(t *testing.T) {
    if findOccurrenceArray(nil)!=nil || findOrderArray("invalid")!=nil || findLastPage(nil)!=0 { t.Fatal("invalid fallback") }
}
func TestPagination(t *testing.T) {
    for _, v:=range []any{"3",float64(3),json.Number("3")} {
        if findLastPage(map[string]any{"meta":map[string]any{"last_page":v}})!=3 {t.Fatal("pagination")}
    }
}
func TestOrders(t *testing.T) {
    if len(findOrderArray(map[string]any{"orders":[]any{1,2}}))!=2 {t.Fatal("orders")}
}
