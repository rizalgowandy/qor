package utils

import (
	"net/url"
	"reflect"
	"testing"
)

func TestParamsMatch(t *testing.T) {
	type paramMatchChecker struct {
		Source      string
		Path        string
		MatchedPath string
		Matched     bool
		Results     url.Values
	}

	checkers := []paramMatchChecker{
		{Source: "", Path: "/", MatchedPath: "/", Results: url.Values{}, Matched: true},
		{Source: "/admin/micro_sites/:id/!preview/", Path: "/admin/micro_sites/1/!preview/index.html", MatchedPath: "/admin/micro_sites/1/!preview/", Results: url.Values{":id": []string{"1"}}, Matched: true},
		{Source: "/hello/:name", Path: "/hello/world", MatchedPath: "/hello/world", Results: url.Values{":name": []string{"world"}}, Matched: true},
		{Source: "/hello/:name/:id", Path: "/hello/world/444", MatchedPath: "/hello/world/444", Results: url.Values{":name": []string{"world"}, ":id": []string{"444"}}, Matched: true},
		{Source: "/hello/:name/:id", Path: "/bye/world/444", MatchedPath: "", Results: nil},
		{Source: "/hello/:name", Path: "/hello/world/444", MatchedPath: "/hello/world", Results: url.Values{":name": []string{"world"}}},
		{Source: "/hello/world", Path: "/hello/name", MatchedPath: "", Results: nil},
		{Source: "/hello/world", Path: "/hello", MatchedPath: "", Results: nil},
		{Source: "/hello/", Path: "/hello", MatchedPath: "/hello", Results: url.Values{}, Matched: true},
		{Source: "/hello/:world", Path: "/hello", MatchedPath: "", Results: nil},
		{Source: "/:locale/campaign", Path: "/en-us/campaign", Matched: true, MatchedPath: "/en-us/campaign", Results: url.Values{":locale": []string{"en-us"}}},
		{Source: "/:locale[(zh|jp)-.*]/campaign", Path: "/zh-CN/campaign", Matched: true, MatchedPath: "/zh-CN/campaign", Results: url.Values{":locale": []string{"zh-CN"}}},
		{Source: `/:locale[(zh|jp)-\w+]/campaign`, Path: "/zh-CN/campaign", Matched: true, MatchedPath: "/zh-CN/campaign", Results: url.Values{":locale": []string{"zh-CN"}}},
		{Source: "/:locale[(zh|jp)-.*]/campaign", Path: "/en-us/campaign", Matched: false, Results: nil},
	}

	for _, checker := range checkers {
		results, matched, ok := ParamsMatch(checker.Source, checker.Path)

		if matched != checker.MatchedPath {
			t.Errorf("%+v's matched path should be %v, but got %v", checker, checker.MatchedPath, matched)
		}

		if ok != checker.Matched {
			t.Errorf("%+v should matched correctly, matched should be %v, but got %v", checker, checker.Matched, ok)
		}

		if !reflect.DeepEqual(results, checker.Results) {
			t.Errorf("%+v's match results should be same, should got %v, but got %v", checker, checker.Results, results)
		}
	}
}
