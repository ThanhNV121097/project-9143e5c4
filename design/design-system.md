# Design System — Hello World Acceptance 8

> Source of truth: approved `index.html`.
> Every value below is extracted from it. Changing a value here without changing approved design is defect.

Last updated: 2026-09-15

## 1. Foundations

### 1.1 Color

Semantic tokens. Name by job, never by hue.

| Token | Value | Used for |
|---|---|---|
| `--color-bg` | `#FFFFFF` | Page background, input background, button text |
| `--color-text` | `#000000` | Body text, heading, input text, input border, focus outline |
| `--color-primary-action` | `#2563EB` | Save button background and border |

#### Contrast audit

Every text-on-background pair actually used. Body text ≥ 4.5:1, large text (≥ 18.66px bold or ≥ 24px) ≥ 3:1, UI borders ≥ 3:1.

| Foreground | Background | Ratio | Passes |
|---|---|---|---|
| `--color-text` | `--color-bg` | `21:1` | AA |
| `--color-bg` | `--color-primary-action` | `5.17:1` | AA |
| `--color-text` | `--color-primary-action` | `4.06:1` | UI focus contrast passes |
| `--color-primary-action` | `--color-bg` | `4.06:1` | UI border contrast passes |

### 1.2 Spacing

Base unit: `4px`. Layout spacing uses this scale unless listed in Known deviations.

| Token | Value |
|---|---|
| `--space-0` | `0` |
| `--space-px` | `1px` |
| `--space-3` | `12px` |
| `--space-5` | `20px` |
| `--space-6` | `24px` |
| `--space-8` | `32px` |
| `--space-10` | `40px` |
| `--space-12` | `48px` |

### 1.3 Typography

Font families:

- Body: `Arial, Helvetica, sans-serif`, loaded from local system fonts.
- Headings: `Arial, Helvetica, sans-serif`, loaded from local system fonts.

| Token | Size | Line height | Weight | Used for |
|---|---|---|---|---|
| `--text-control` | `18px` | `normal` | `400` | Input text |
| `--text-button` | `18px` | `normal` | `700` | Save button |
| `--text-heading` | `clamp(40px, 9vw, 80px)` | `1.05` | `700` | h1 greeting |

Heading levels are used in order and never skipped for visual sizing.

| Token | Value | Used for |
|---|---|---|
| `--font-weight-body` | `400` | Running text and input |
| `--font-weight-heading` | `700` | Greeting heading |
| `--font-weight-action` | `700` | Save button |
| `--tracking-tight-heading` | `-0.04em` | Greeting heading |

### 1.4 Radius, border, shadow, motion

| Token | Value | Used for |
|---|---|---|
| `--radius-square` | `0` | Input and button |
| `--border-width` | `1px` | Input and button border |
| `--focus-ring-width` | `3px` | Input and button focus outline |
| `--focus-ring-offset` | `3px` | Input and button focus outline offset |

Motion: no animation or transition appears in approved design.

### 1.5 Layout and breakpoints

| Name | Max width | Container | Columns | Gutter |
|---|---|---|---|---|
| `mobile-stack` | `520px` | `min(100%, 560px)` | 1 | `12px` |
| `default` | Above `520px` | `min(100%, 560px)` | 1 section; form row | `12px` |

| Layer | Value |
|---|---|
| Base | `0` |

## 2. Components

### 2.1 Greeting section

**Purpose** — Center single greeting task on page. Use for only one primary greeting heading with update form.

**Anatomy** — `[h1 greeting] [form: label input save button]`.

**Variants**

| Variant | Tokens | When to use |
|---|---|---|
| Default | `--color-bg`, `--color-text`, `--space-6`, `--space-8`, `--text-heading` | Only page section |

**Sizes**

| Size | Width | Min height | Text token |
|---|---|---|---|
| Default | `min(100%, 560px)` | `calc(100vh - 48px)` | `--text-heading` |
| Mobile | `min(100%, 560px)` | `calc(100vh - 40px)` | `--text-heading` |

**States**

| State | Visual change | Tokens |
|---|---|---|
| Default | Centered heading above form, black text on white background | `--color-bg`, `--color-text` |

**Accessibility** — Section uses `aria-labelledby` pointing at h1. Greeting text remains text, not image. Heading may wrap anywhere for long saved greetings.

### 2.2 Greeting input

**Purpose** — Edit stored greeting. Do not use for search or multiline text.

**Anatomy** — `[visually hidden label] [text input]`.

**Variants**

| Variant | Tokens | When to use |
|---|---|---|
| Default | `--color-bg`, `--color-text`, `--border-width`, `--radius-square`, `--text-control` | Single-line greeting edit |

**Sizes**

| Size | Height | Padding | Text token |
|---|---|---|---|
| Default | `48px` | `0 14px` | `--text-control` |

**States**

| State | Visual change | Tokens |
|---|---|---|
| Default | White field, black text, black 1px border | `--color-bg`, `--color-text`, `--border-width` |
| Focus | Black 3px outline with 3px offset | `--color-text`, `--focus-ring-width`, `--focus-ring-offset` |

**Accessibility** — Uses explicit label hidden visually. `required` prevents empty browser submission. Keyboard focus indicator visible. Minimum hit target is `48px` tall.

### 2.3 Save button

**Purpose** — Submit edited greeting. Use only for primary save action on this page.

**Anatomy** — `[label]`.

**Variants**

| Variant | Tokens | When to use |
|---|---|---|
| Primary action | `--color-primary-action`, `--color-bg`, `--border-width`, `--radius-square`, `--text-button` | Save greeting |

**Sizes**

| Size | Height | Padding | Text token |
|---|---|---|---|
| Default | `48px` | `0 22px` | `--text-button` |

**States**

| State | Visual change | Tokens |
|---|---|---|
| Default | Blue background and border, white bold label | `--color-primary-action`, `--color-bg`, `--text-button` |
| Hover | Same as default | `--color-primary-action`, `--color-bg` |
| Focus | Black 3px outline with 3px offset | `--color-text`, `--focus-ring-width`, `--focus-ring-offset` |

**Accessibility** — Native `button type="submit"`; Enter key submits form. Keyboard focus indicator visible. Minimum hit target is `48px` tall.

## 3. Content and formatting

- Voice and tone: plain, direct, minimal.
- Date, time, number, and currency formats: not shown in approved design.
- Capitalization rule: greeting preserves stored text exactly; button uses title case label `Save`; label uses title case `Greeting`.
- Empty-state and error-message wording pattern: no visible empty or error state in approved design; empty submit refocuses input without message.

## 4. Known deviations

| Where | Deviation | Why it stands | Follow-up |
|---|---|---|---|
| `.greeting-input` padding | `14px` is not on 4px spacing scale | Approved design uses exact value | Keep unless future design revision changes input padding |
| Form validation | Empty submit has no visible error message | Approved design draws no error state | Add visible error only if stakeholder requests error state |

## 5. Change log

| Date | Change | Design PR |
|---|---|---|
| 2026-09-15 | Initial design system extracted from approved `index.html` | This PR |
