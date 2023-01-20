package usecase

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/Anton-Hudz/MovieList/app/entities"
)

func createQuery(params entities.QueryParams) (string, error) {
	var (
		query        string = `SELECT f.id, f.name, f.genre, d.name, f.rate, f.year, f.minutes FROM films f JOIN directors d ON f.director_id = d.id `
		SQL          string
		defaultLimit int = 5
	)

	if params.Genre != "" {
		splitgenre := strings.Split(params.Genre, ",")
		for i := range splitgenre {
			splitgenre[i] = fmt.Sprintf("'%s'", splitgenre[i])
		}
		newstr := strings.Join(splitgenre, ",")
		SQL = fmt.Sprintf("%sWHERE genre IN (%s) ", query, newstr)
	}

	if params.Genre != "" && params.Rate != "" {
		SQL = fmt.Sprintf("%sAND ", SQL)
	}

	if params.Genre == "" && params.Rate != "" {
		SQL = fmt.Sprintf("%sWHERE ", query)
	}

	if params.Rate != "" {
		splitrange := strings.Split(params.Rate, "-")
		if len(splitrange) != 2 {
			return "", errors.New("error rate must consist of min & max numbers")
		}
		for i := range splitrange {
			_, err := strconv.ParseFloat(splitrange[i], 32)
			if err != nil {
				return "", errors.New("error rate parameters must be a numbers")
			}
		}
		SQL = fmt.Sprintf("%s(rate >= %s AND rate <= %s) ", SQL, splitrange[0], splitrange[1])
	}

	if params.Genre == "" && params.Rate == "" {
		SQL = query
	}

	if params.Sort != "" {
		condition := strings.ReplaceAll(params.Sort, ",", ", ")
		sortString := fmt.Sprintf("ORDER BY %s ", condition)
		SQL = SQL + sortString
	}

	if params.Limit != "" {
		limit, err := strconv.Atoi(params.Limit)
		if err != nil {
			return "", errors.New("error limit must be number")
		}
		if limit == 0 {
			limit = defaultLimit
		}
		limitString := fmt.Sprintf("LIMIT %s ", strconv.Itoa(limit))
		SQL = SQL + limitString

	}
	if params.Limit == "" {
		limitString := fmt.Sprintf("LIMIT %s ", strconv.Itoa(defaultLimit))
		SQL = SQL + limitString
	}

	if params.Offset != "" {
		offset, err := strconv.Atoi(params.Offset)
		if err != nil {
			return "", errors.New("error offset must be number")
		}
		offsetString := fmt.Sprintf("OFFSET %s;", strconv.Itoa(offset))
		SQL = SQL + offsetString
	}

	return SQL, nil
}
