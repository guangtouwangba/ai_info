# When Detection Fails: The Power of Fine-Tuned Models to Generate Human-Like Social Media Text

**URL**: http://arxiv.org/abs/2506.09975v1

## 原始摘要

Detecting AI-generated text is a difficult problem to begin with; detecting
AI-generated text on social media is made even more difficult due to the short
text length and informal, idiosyncratic language of the internet. It is
nonetheless important to tackle this problem, as social media represents a
significant attack vector in online influence campaigns, which may be bolstered
through the use of mass-produced AI-generated posts supporting (or opposing)
particular policies, decisions, or events. We approach this problem with the
mindset and resources of a reasonably sophisticated threat actor, and create a
dataset of 505,159 AI-generated social media posts from a combination of
open-source, closed-source, and fine-tuned LLMs, covering 11 different
controversial topics. We show that while the posts can be detected under
typical research assumptions about knowledge of and access to the generating
models, under the more realistic assumption that an attacker will not release
their fine-tuned model to the public, detectability drops dramatically. This
result is confirmed with a human study. Ablation experiments highlight the
vulnerability of various detection algorithms to fine-tuned LLMs. This result
has implications across all detection domains, since fine-tuning is a generally
applicable and realistic LLM use case.


## AI 摘要

检测社交媒体上的AI生成文本极具挑战性，主要因为文本短且语言非正式。研究者模拟威胁者视角，创建了50.5万条AI生成的社交媒体帖子，涵盖11个争议话题。研究发现，在已知生成模型的理想条件下可检测AI文本，但现实中攻击者不会公开其微调模型，检测效果大幅下降。人类实验证实了这一结论。消融实验显示，现有检测算法对微调模型普遍脆弱。这一发现对所有检测领域都有影响，因为模型微调是LLM的常见应用场景。研究突显了AI生成内容检测在真实场景中的局限性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-12T14:01:52Z
- **目录日期**: 2025-06-12
