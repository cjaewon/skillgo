# skillgo
![Kakao Skill Payload Version](http://img.shields.io/badge/Kakao%20SkillPayload%20Version-v2-yellow?style=flat-square&)

📦 Go 을 위한 kakao i open builder skill 응답 포맷 생성 모듈  
[🟡 JavaScript Version](https://github.com/cjaewon/kakaoEmbed)

## 지원하는 SkillResponse

### 기본

| SkillResponse | 타입 (struct) | 함수 이름 |
| - | - | - |
| **SimpleText** | `SimpleTextType` | SimpleText |
| **SimpleImage** | `SimpleImageType` | SimpleImage |
| **BasicCard** | `BasicCardType` | BasicCard |
| **CommerceCard** | `CommerceCardType` | CommerceCard |
| **ListCard** | `ListCard` | ListCard

### 공통
[kakao i 오픈빌더 공통](https://i.kakao.com/docs/skill-response-format#%EA%B3%B5%ED%86%B5)

| SkillResponse | 타입 (struct) | 함수 이름 |
| - | - | - |
| **Thumbnail** | `ThumbnailType` | Thumbnail |
| **Link** | `LinkType` | Link |
| **Button** | `ButtonType` | Button |
| **Profile** | `ProfileType` | Profile |
