# Tracing Thought: Using Chain-of-Thought Reasoning to Identify the LLM Behind AI-Generated Text

**URL**: http://arxiv.org/abs/2504.16913v1

## 原始摘要

In recent years, the detection of AI-generated text has become a critical
area of research due to concerns about academic integrity, misinformation, and
ethical AI deployment. This paper presents COT Fine-tuned, a novel framework
for detecting AI-generated text and identifying the specific language model.
responsible for generating the text. We propose a dual-task approach, where
Task A involves classifying text as AI-generated or human-written, and Task B
identifies the specific LLM behind the text. The key innovation of our method
lies in the use of Chain-of-Thought reasoning, which enables the model to
generate explanations for its predictions, enhancing transparency and
interpretability. Our experiments demonstrate that COT Fine-tuned achieves high
accuracy in both tasks, with strong performance in LLM identification and
human-AI classification. We also show that the CoT reasoning process
contributes significantly to the models effectiveness and interpretability.


## AI 摘要

该论文提出了一种名为COT Fine-tuned的新型AI文本检测框架，通过双任务设计同时实现AI生成文本检测（任务A）和识别具体语言模型（任务B）。其核心创新在于采用思维链推理机制，使模型能生成预测解释，从而提升透明度和可解释性。实验表明，该方法在两项任务中均取得高准确率，尤其在特定LLM识别方面表现突出。研究证实思维链推理过程对模型效能和可解释性具有显著贡献，为解决学术诚信、错误信息和AI伦理部署等问题提供了有效技术方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-24T05:01:23Z
- **目录日期**: 2025-04-24
