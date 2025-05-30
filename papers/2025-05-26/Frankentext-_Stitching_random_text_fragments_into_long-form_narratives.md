# Frankentext: Stitching random text fragments into long-form narratives

**URL**: http://arxiv.org/abs/2505.18128v1

## 原始摘要

We introduce Frankentexts, a new type of long-form narratives produced by
LLMs under the extreme constraint that most tokens (e.g., 90%) must be copied
verbatim from human writings. This task presents a challenging test of
controllable generation, requiring models to satisfy a writing prompt,
integrate disparate text fragments, and still produce a coherent narrative. To
generate Frankentexts, we instruct the model to produce a draft by selecting
and combining human-written passages, then iteratively revise the draft while
maintaining a user-specified copy ratio. We evaluate the resulting Frankentexts
along three axes: writing quality, instruction adherence, and detectability.
Gemini-2.5-Pro performs surprisingly well on this task: 81% of its Frankentexts
are coherent and 100% relevant to the prompt. Notably, up to 59% of these
outputs are misclassified as human-written by detectors like Pangram, revealing
limitations in AI text detectors. Human annotators can sometimes identify
Frankentexts through their abrupt tone shifts and inconsistent grammar between
segments, especially in longer generations. Beyond presenting a challenging
generation task, Frankentexts invite discussion on building effective detectors
for this new grey zone of authorship, provide training data for mixed
authorship detection, and serve as a sandbox for studying human-AI co-writing
processes.


## AI 摘要

研究人员提出"Frankentexts"——一种由大语言模型生成的文本，要求90%内容直接复制自人类作品。该任务测试模型在严格限制下的可控生成能力，需整合碎片化文本并保持连贯性。实验显示，Gemini-2.5-Pro表现优异：81%生成文本连贯且符合提示，59%被检测器误判为人类创作。人类可通过语气突变和语法不一致识别这类文本。该研究揭示了AI检测器的局限性，为混合创作检测提供训练数据，并为研究人机协作写作建立了实验平台，探讨了新型"灰色地带"文本的检测方法。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-26T23:01:47Z
- **目录日期**: 2025-05-26
