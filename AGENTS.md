# AGENTS.md

## 커밋 메시지 규칙

- 커밋 메시지는 한국어로 작성한다.
- 제목은 변경의 핵심을 간결하게 요약한다.
- 본문은 반드시 작성하며, 무엇을 왜 했는지 의도가 드러나도록 한다.
- 본문은 불릿 리스트 대신 문단 형태로 작성한다.

## 디렉터리 구조

- `sandbox.md` : `update-sandbox.sh`로 매일 자동 수집되는 개발자 트렌드 raw 피드. 파일을 분할하지 않고 단일 파일에 계속 누적한다.
- `summaries/YYYY-MM.md` : 월간 합성 요약. Phase 2의 `/synthesize-month` 스킬이 초안을 만들고 사용자가 확정한다.
- `notes/` : 수동 작성 노트와 외부 글 번역 등 evergreen 콘텐츠. `notes/translations/`에 `analyze-article` 스킬 산출물이 모인다.
- `archive/` : dormant하지만 보존 가치 있는 과거 기록.
- 파일명 슬러그는 영문 kebab-case를 기본으로 하되, 이미 존재하는 한국어 파일명은 이관 시 그대로 유지한다.
