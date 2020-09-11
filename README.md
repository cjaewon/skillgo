# skillgo
![Kakao Skill Payload Version](http://img.shields.io/badge/Kakao%20SkillPayload%20Version-v2-yellow?style=flat-square&)
[![Go Report Card](https://goreportcard.com/badge/github.com/cjaewon/skillgo)](https://goreportcard.com/report/github.com/cjaewon/skillgo)

📦 Go 을 위한 kakao i open builder skill 응답 포맷 생성 모듈  

[🗞️ Docs - pkg.go.dev](https://pkg.go.dev/github.com/cjaewon/skillgo?tab=doc)  
[🟡 JavaScript Version](https://github.com/cjaewon/kakaoEmbed)  

## 설치
```sh
go get github.com/cjaewon/skillgo
```

## 주요 타입
| 이름 | 타입 |
| - | - |
| [**SkillPayload**](https://i.kakao.com/docs/skill-response-format#skillpayload) | `SkillPayload` |
| [**SkillResponse**](https://i.kakao.com/docs/skill-response-format#skillresponse) | `SkillResponse` |
| [**SkillTemplate**](https://i.kakao.com/docs/skill-response-format#skilltemplate) | `SkillTemplate` |
| [**ContextControl**](https://i.kakao.com/docs/skill-response-format#contextcontrol) | `ContextControl` |
| [**QuickReplies**](https://i.kakao.com/docs/skill-response-format#quickreplies) | `QuickReplies` |

## 지원하는 SkillResponse

### 기본

| SkillResponse | 타입 (struct) | 함수 이름 |
| - | - | - |
| [**SimpleText**](https://i.kakao.com/docs/skill-response-format#simpletext) | `SimpleTextType` | SimpleText |
| [**SimpleImage**](https://i.kakao.com/docs/skill-response-format#simpleimage) | `SimpleImageType` | SimpleImage |
| [**BasicCard**](https://i.kakao.com/docs/skill-response-format#basiccard) | `BasicCardType` | BasicCard |
| [**CommerceCard**](https://i.kakao.com/docs/skill-response-format#commercecard) | `CommerceCardType` | CommerceCard |
| [**ListCard**](https://i.kakao.com/docs/skill-response-format#listcard) | `ListCardType` | ListCard |
| [**Carousel**](https://i.kakao.com/docs/skill-response-format#carousel) | `CarouselType` | Carousel |

### 공통
[kakao i 오픈빌더 공통](https://i.kakao.com/docs/skill-response-format#%EA%B3%B5%ED%86%B5)

| SkillResponse | 타입 (struct) | 함수 이름 |
| - | - | - |
| [**Thumbnail**](https://i.kakao.com/docs/skill-response-format#thumbnail) | `ThumbnailType` | Thumbnail |
| [**Link**](https://i.kakao.com/docs/skill-response-format#link) | `LinkType` | Link |
| [**Button**](https://i.kakao.com/docs/skill-response-format#button) | `ButtonType` | Button |
| [**Profile**](https://i.kakao.com/docs/skill-response-format#profile) | `ProfileType` | Profile |
| [**CarouselHeader**](https://i.kakao.com/docs/skill-response-format#carouselheader) | `CarouselHeaderType` | CarouselHeader |

## 예제

[📄 **Simple (SimpleText, SimpleImage)**](https://github.com/cjaewon/skillgo/tree/master/_examples/simple)  
[📄 **Card (BasicCard, CommerceCard, Listcard)**](https://github.com/cjaewon/skillgo/tree/master/_examples/card)  
