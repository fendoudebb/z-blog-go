package repo

import (
	"fmt"
	"z-blog-go/db"
)

func GetSitemap() (string, error) {
	baseUrl := `http://localhost:8080`
	sql := fmt.Sprintf(`
select
xmlroot(
	xmlelement(name urlset, xmlattributes('http://www.sitemaps.org/schemas/sitemap/0.9' as xmlns),
		xmlconcat(
			xmlelement(name url, 
				xmlelement(name loc, '%s'),
				xmlelement(name lastmod, current_date),
				xmlelement(name changefreq, 'always'),
				xmlelement(name priority, 1)
			),
			xmlagg(
				xmlelement(name url,
					xmlforest(concat('%s/p/',id,'.html') as loc),
					xmlelement(name lastmod, current_date),
					xmlelement(name changefreq, 'daily'),
					xmlelement(name priority, 0.8)
				)
			order by id desc)
		)
	)
	,version '1.0', standalone yes)::text as sitemap
from post where post_status=0
`, baseUrl, baseUrl)

	row := db.DB.QueryRow(sql)

	var sitemap string
	if err := row.Scan(&sitemap); err != nil {
		return sitemap, fmt.Errorf("GetSitemap: %v", err)
	}
	return sitemap, nil
}
