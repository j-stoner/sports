package mlbmock

import (
	"context"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"time"

	yaml "github.com/ghodss/yaml"
	"github.com/markbates/pkger"
	"github.com/robbydyer/sports/pkg/logo"
	"github.com/robbydyer/sports/pkg/mlb"
	"github.com/robbydyer/sports/pkg/sportboard"
	"github.com/robbydyer/sports/pkg/util"
)

type MockMLBAPI struct {
	teams           []*mlb.Team
	games           map[string][]*mlb.Game
	logos           map[string]*logo.Logo
	logoSourceCache map[string]image.Image
	log             *log.Logger
}

func (m *MockMLBAPI) GetTeams(ctx context.Context) ([]sportboard.Team, error) {
	var tList []sportboard.Team

	for _, t := range m.teams {
		tList = append(tList, t)
	}

	return tList, nil
}
func (m *MockMLBAPI) GetScheduledGames(ctx context.Context, date time.Time) ([]sportboard.Game, error) {
	dateStr := m.DateStr(date)
	var gList []sportboard.Game

	for _, g := range m.games[dateStr] {
		gList = append(gList, g)
	}

	return gList, nil
}
func (m *MockMLBAPI) DateStr(d time.Time) string {
	return d.Format(mlb.DateFormat)
}
func (m *MockMLBAPI) League() string {
	return "Fake MLB"
}
func (m *MockMLBAPI) GetLogo(logoKey string, logoConf *logo.Config, bounds image.Rectangle) (*logo.Logo, error) {
	fullLogoKey := fmt.Sprintf("%s_%dx%d", logoKey, bounds.Dx(), bounds.Dy())
	l, ok := m.logos[fullLogoKey]
	if ok {
		return l, nil
	}

	sources, err := m.logoSources()
	if err != nil {
		return nil, err
	}

	l, err = mlb.GetLogo(logoKey, logoConf, bounds, sources)
	if err != nil {
		return nil, err
	}

	m.logos[fullLogoKey] = l

	return l, nil
}
func (m *MockMLBAPI) logoSources() (map[string]image.Image, error) {

	if len(m.logoSourceCache) == len(mlb.ALL) {
		return m.logoSourceCache, nil
	}

	for _, t := range mlb.ALL {
		f, err := pkger.Open(fmt.Sprintf("github.com/robbydyer/sports:/pkg/mlb/assets/logos/%s.png", t))
		if err != nil {
			return nil, fmt.Errorf("failed to locate logo asset: %w", err)
		}
		defer f.Close()

		i, err := png.Decode(f)
		if err != nil {
			return nil, err
		}

		m.logoSourceCache[t] = i
	}

	return m.logoSourceCache, nil
}
func (m *MockMLBAPI) AllTeamAbbreviations() []string {
	return mlb.ALL
}

func (m *MockMLBAPI) UpdateTeams(ctx context.Context) error {
	return nil
}
func (m *MockMLBAPI) UpdateGames(ctx context.Context, dateStr string) error {
	return nil
}
func (m *MockMLBAPI) TeamFromAbbreviation(ctx context.Context, abbrev string) (sportboard.Team, error) {
	for _, t := range m.teams {
		if t.Abbreviation == abbrev {
			return t, nil
		}
	}

	return nil, fmt.Errorf("could not find team with abbreviation '%s'", abbrev)
}

func MockLiveGameGetter(ctx context.Context, link string) (sportboard.Game, error) {
	f, err := pkger.Open("github.com/robbydyer/sports:/pkg/mlbmock/assets/mock_livegames.yaml")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	dat, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var gameList []*mlb.Game

	if err := yaml.Unmarshal(dat, &gameList); err != nil {
		return nil, fmt.Errorf("failed to unmarshal live game mock yaml: %w", err)
	}

	for _, liveGame := range gameList {
		if liveGame.Link == link {
			liveGame.GameTime = time.Now().Local()
			return liveGame, nil
		}
	}

	return nil, fmt.Errorf("could not locate live game with Link '%s'", link)
}

func New() (*MockMLBAPI, error) {
	// Load Teams
	f, err := pkger.Open("github.com/robbydyer/sports:/pkg/mlbmock/assets/mock_teams.yaml")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	dat, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var teamList []*mlb.Team

	if err := yaml.Unmarshal(dat, &teamList); err != nil {
		return nil, err
	}

	// Load Games
	gamef, err := pkger.Open("github.com/robbydyer/sports:/pkg/mlbmock/assets/mock_games.yaml")
	if err != nil {
		return nil, err
	}
	defer gamef.Close()

	dat, err = ioutil.ReadAll(gamef)
	if err != nil {
		return nil, err
	}

	var gameList []*mlb.Game

	if err := yaml.Unmarshal(dat, &gameList); err != nil {
		return nil, fmt.Errorf("failed to unmarshal mock yaml: %w", err)
	}

	for _, g := range gameList {
		g.GameGetter = MockLiveGameGetter
	}

	today := util.Today().Format(mlb.DateFormat)
	m := &MockMLBAPI{
		games: map[string][]*mlb.Game{
			today: gameList,
		},
		teams:           teamList,
		logos:           make(map[string]*logo.Logo),
		logoSourceCache: make(map[string]image.Image),
	}

	return m, nil
}
