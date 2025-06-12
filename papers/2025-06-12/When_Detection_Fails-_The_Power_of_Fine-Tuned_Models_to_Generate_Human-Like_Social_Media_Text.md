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

检测社交媒体上的AI生成文本极具挑战性，主要由于文本短小且语言非正式。研究者模拟威胁行为者，创建了包含505,159条AI生成帖子的数据集，涵盖11个争议话题。研究发现，在已知生成模型的理想条件下可有效检测，但现实中攻击者不会公开其微调模型，导致检测率大幅下降。人类实验证实了这一结果，且不同检测算法对微调模型均表现脆弱。这一发现对AI文本检测领域具有广泛影响，因为模型微调是常见且实际的应用场景。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-12T17:01:57Z
- **目录日期**: 2025-06-12
