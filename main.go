package main

import (
	"time"
)

type Data struct {
	a, b, c, d, e, f, g DataValue
}

type DataPointer struct {
	a, b, c, d, e, f, g *DataValue
}

type DataValue struct {
	Type, Value string
}

func NewData() Data {
	data := Data{}
	data.a = DataValue{
		Type:  "a",
		Value: "a",
	}
	data.b = DataValue{
		Type:  "b",
		Value: "b",
	}
	data.c = DataValue{
		Type:  "c",
		Value: "c",
	}
	data.d = DataValue{
		Type:  "d",
		Value: "d",
	}
	data.e = DataValue{
		Type:  "e",
		Value: "e",
	}
	data.f = DataValue{
		Type:  "f",
		Value: "f",
	}
	data.g = DataValue{
		Type:  "g",
		Value: "g",
	}

	return data
}

func NewDataPointer() DataPointer {
	data := DataPointer{}
	data.a = &DataValue{
		Type:  "a",
		Value: "a",
	}
	data.b = &DataValue{
		Type:  "b",
		Value: "b",
	}
	data.c = &DataValue{
		Type:  "c",
		Value: "c",
	}
	data.d = &DataValue{
		Type:  "d",
		Value: "d",
	}
	data.e = &DataValue{
		Type:  "e",
		Value: "e",
	}
	data.f = &DataValue{
		Type:  "f",
		Value: "f",
	}
	data.g = &DataValue{
		Type:  "g",
		Value: "g",
	}

	return data
}

func SummarizeData(data Data) string {
	summarized := ""

	summarized += data.a.Value
	summarized += data.b.Value
	summarized += data.c.Value
	summarized += data.d.Value
	summarized += data.e.Value
	summarized += data.f.Value
	summarized += data.g.Value

	return summarized
}

func SummarizeDataPointer(data DataPointer) string {
	summarized := ""

	summarized += data.a.Value
	summarized += data.b.Value
	summarized += data.c.Value
	summarized += data.d.Value
	summarized += data.e.Value
	summarized += data.f.Value
	summarized += data.g.Value

	return summarized
}

func main() {
	completeSummary := make([]string, 0)
	ini := time.Now()

	for range 100000 {
		completeSummary = append(completeSummary, SummarizeData(NewData()))
	}
	println("Time to summarize concrete data: ", time.Since(ini).Microseconds())
	//Time to summarize concrete data:  16778

	completeSummary = make([]string, 0)
	ini = time.Now()

	for range 1000000 {
		completeSummary = append(completeSummary, SummarizeDataPointer(NewDataPointer()))
	}
	println("Time to summarize pointer data: ", time.Since(ini).Microseconds())
	//Time to summarize pointer data:  150523
}
