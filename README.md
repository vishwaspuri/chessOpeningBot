The ECO Codes is an Encyclopedia classification system for the chess openings moves. It is a repository of the most important or the Top 100 chess opening moves compiled by an organization called Chess Informant.

The REST API provided can be used to fetch any of these openings by their code or all at once.

### API GUIDE

---
```
    GET / 
    Response:
    {
        "data": [
            {
                "code": "A38",
                "creator": "English, Symmetrical",
                "opening": "1 c4 c5 2 Nc3 Nc6 3 g3 g6 4 Bg2 Bg7 5 Nf3 Nf6"
            },
            ...
        ]
    }
```
```
    GET /<eco-code>  
    Example: GET /C87
    Response:
    {
        "data": {
            "code": "C87",
            "creator": "Ruy Lopez",
            "opening": "1 e4 e5 2 Nf3 Nc6 3 Bb5 a6 4 Ba4 Nf6 5 O-O Be7 6 Re1 d6"
        }
    }

```
