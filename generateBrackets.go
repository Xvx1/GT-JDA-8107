package main

import (
    "fmt"
    "sort"
    "flag"
)

func main() {
    Alex := Player{0, "Alex", 4}
    Daniel := Player{1, "Daniel", 3}
    Deven  := Player{2, "Deven", 2}
    Dennis := Player{3, "Dennis", 1}
    Moss := Player{4, "Dr. Moss", 111}
    Madison := Player{5, "Madison", -999999}
    Arjun := Player{6, "Arjun", 1111}
    Ben := Player{7, "Ben", 3}
    players := []*Player{&Alex, &Daniel, &Deven, &Dennis, &Moss, &Madison, &Arjun, &Ben}
    
    single := flag.Bool("single", false, "Single Elimination")
    roundrobin := flag.Int("roundrobin", -1, "Number of Round Robin Groups")
    flag.Parse()

    if *single {
        fmt.Println("-----------------------------------------")
        fmt.Println("| Single Elimination Tournament Bracket |")
        fmt.Println("-----------------------------------------")
        singleElimBracket := generateSingleElimBrackets(players)
        for i := 0; i < len(singleElimBracket); i++ {
            for j := 0; j < len(singleElimBracket[i]); j++ {
                A := singleElimBracket[i][j].A.String()
                B := singleElimBracket[i][j].B.String()
                if !singleElimBracket[i][j].A.IsPlayer() {
                    A = "Winner of " + A
                }
                if !singleElimBracket[i][j].B.IsPlayer() {
                    B = "Winner of " + B
                }
                fmt.Println("Match " + fmt.Sprintf("%d", singleElimBracket[i][j].Number) + ": " + A + " vs " + B)
            }
        }
    }
    
    if *roundrobin > 0 {
        fmt.Println("-----------------------------------------")
        fmt.Println("|           Round Robin Groups          |")
        fmt.Println("-----------------------------------------")
        roundRobinGroups := generateRoundRobinGroups(players, *roundrobin)
        for groupNumber, group := range roundRobinGroups {
            fmt.Println("Group " + fmt.Sprintf("%d:", groupNumber + 1))
            for _, player := range group {
                fmt.Println(player.String())
            }
            fmt.Println()
        }
    }
}

type Match interface {
    String() string
    IsPlayer() bool
}

type Matchup struct {
    A Match
    B Match
    Seed int
    Number int
}

func (m Matchup) String() (string) {
    return "Match " + fmt.Sprintf("%d", m.Number)
}

func (m Matchup) IsPlayer() (bool) {
    return false
}

type Player struct {
    Id int
    Name string
    Seed int
}

func (p Player) String() (string) {
    return p.Name
}

func (p Player) IsPlayer() (bool) {
    return true
}

type BySeedMatchup []*Matchup

func (s BySeedMatchup) Len() (int) {
    return len(s)
}

func (s BySeedMatchup) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s BySeedMatchup) Less(i, j int) (bool) {
    return s[i].Seed < s[j].Seed
}

type BySeedPlayer []*Player

func (s BySeedPlayer) Len() (int) {
    return len(s)
}

func (s BySeedPlayer) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s BySeedPlayer) Less(i, j int) (bool) {
    return s[i].Seed < s[j].Seed
}

func isPowerOfTwo(i int) (bool) {
    return (i > 0) && ((i & (i - 1)) == 0)
}

func generateSingleElimBrackets(playerList []*Player) ([][]*Matchup) {
    players := append([]*Player(nil), playerList...)
    sort.Stable(sort.Reverse(BySeedPlayer(players)))
    nilCounter := 0
    for ; !isPowerOfTwo(len(players)); {
        playerList = append(players, nil)
        nilCounter++
    }
    numPlayers := len(players)
    matchCounter := 1
    firstMatchLayer := generateMatchLayerFromPlayers(players, matchCounter)
    matchCounter += len(players) / 2
    matchLayers := [][]*Matchup{firstMatchLayer}
    currMatchLayer := firstMatchLayer
    for ; len(currMatchLayer) != 1; {
        currMatchLayer = generateMatchLayerFromMatchups(currMatchLayer, matchCounter)
        matchLayers = append(matchLayers, currMatchLayer)
        matchCounter += len(currMatchLayer)
    }
    removeNilMatchups(&matchLayers, numPlayers, nilCounter)
    return matchLayers
}

func generateMatchLayerFromPlayers(matches []*Player, number int) ([]*Matchup) {
    var nextMatchLayer []*Matchup
    for i := 0; i < len(matches) / 2; i++ {
        a := new(Matchup)
        (*a).A = matches[i]
        (*a).B = matches[len(matches) - (i + 1)]
        (*a).Seed = matches[i].Seed
        (*a).Number = number + i
        nextMatchLayer = append(nextMatchLayer, a)
    }
    return nextMatchLayer
}

func generateMatchLayerFromMatchups(matches []*Matchup, number int) ([]*Matchup) {
    sort.Stable(sort.Reverse(BySeedMatchup(matches)))
    var nextMatchLayer []*Matchup
    for i := 0; i < len(matches) / 2; i++ {
        a := new(Matchup)
        (*a).A = matches[i]
        (*a).B = matches[len(matches) - (i + 1)]
        (*a).Seed = matches[i].Seed
        (*a).Number = number + i
        nextMatchLayer = append(nextMatchLayer, a)
    }
    return nextMatchLayer
}

func removeNilMatchups(bracket *[][]*Matchup, numPlayers int, numNils int) {
    matchNumber := numPlayers - numNils - 1
    for i := len(*bracket) - 1; i >= 0; i-- {
        for j := len((*bracket)[i]) - 1; j >= 0; j-- {
            if (*bracket)[i][j].A.IsPlayer() {
                if (*bracket)[i][j].B.(*Player) != nil {
                    (*bracket)[i][j].Number = matchNumber
                    matchNumber--
                } else {
                    copy((*bracket)[i][j:], (*bracket)[i][j+1:])
                    (*bracket)[i][len((*bracket)[i])-1] = nil
                    (*bracket)[i] = (*bracket)[i][:len((*bracket)[i])-1]
                }
            } else {
                retrieveMatchup((*bracket)[i][j], matchNumber)
                matchNumber--
            }
        }
    }
}

func retrieveMatchup(match *Matchup, number int) {
    match.Number = number
    if match.A.(*Matchup).A.IsPlayer() && match.A.(*Matchup).B.(*Player) == nil {
        match.A = match.A.(*Matchup).A.(*Player)
    }
    if match.B.(*Matchup).A.IsPlayer() && match.B.(*Matchup).B.(*Player) == nil {
        match.B = match.B.(*Matchup).A.(*Player)
    }
}

func generateRoundRobinGroups(playerList []*Player, numGroups int) ([][]*Player) {
    players := append([]*Player(nil), playerList...)
    sort.Stable(sort.Reverse(BySeedPlayer(players)))
    for ; len(players) % numGroups != 0; {
        players = append(players, nil)
    }
    groups := make([][]*Player, numGroups)
    for i := 0; i < numGroups; i++ {
        for j := i; j < len(players); j += numGroups {
            if players[j] != nil {
                groups[i] = append(groups[i], players[j])
            }
        }
    }
    return groups
}