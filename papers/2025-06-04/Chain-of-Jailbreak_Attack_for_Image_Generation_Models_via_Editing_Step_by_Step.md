# Chain-of-Jailbreak Attack for Image Generation Models via Editing Step by Step

**URL**: http://arxiv.org/abs/2410.03869v2

## 原始摘要

Text-based image generation models, such as Stable Diffusion and DALL-E 3,
hold significant potential in content creation and publishing workflows, making
them the focus in recent years. Despite their remarkable capability to generate
diverse and vivid images, considerable efforts are being made to prevent the
generation of harmful content, such as abusive, violent, or pornographic
material. To assess the safety of existing models, we introduce a novel
jailbreaking method called Chain-of-Jailbreak (CoJ) attack, which compromises
image generation models through a step-by-step editing process. Specifically,
for malicious queries that cannot bypass the safeguards with a single prompt,
we intentionally decompose the query into multiple sub-queries. The image
generation models are then prompted to generate and iteratively edit images
based on these sub-queries. To evaluate the effectiveness of our CoJ attack
method, we constructed a comprehensive dataset, CoJ-Bench, encompassing nine
safety scenarios, three types of editing operations, and three editing
elements. Experiments on four widely-used image generation services provided by
GPT-4V, GPT-4o, Gemini 1.5 and Gemini 1.5 Pro, demonstrate that our CoJ attack
method can successfully bypass the safeguards of models for over 60% cases,
which significantly outperforms other jailbreaking methods (i.e., 14%).
Further, to enhance these models' safety against our CoJ attack method, we also
propose an effective prompting-based method, Think Twice Prompting, that can
successfully defend over 95% of CoJ attack. We release our dataset and code to
facilitate the AI safety research.


## AI 摘要

研究人员提出了一种针对文本生成图像模型（如Stable Diffusion、DALL-E 3等）的新型越狱攻击方法"Chain-of-Jailbreak"（CoJ）。该方法通过将恶意查询分解为多个子查询，引导模型分步生成并迭代编辑图像，成功绕过了现有安全防护。实验显示，CoJ对GPT-4V、Gemini等主流模型的攻击成功率超过60%，远超传统方法（14%）。为应对此威胁，研究团队还开发了"Think Twice Prompting"防御策略，可有效拦截95%以上的CoJ攻击。相关数据集和代码已开源，以促进AI安全研究发展。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-04T23:02:26Z
- **目录日期**: 2025-06-04
