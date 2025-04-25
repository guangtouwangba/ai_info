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

这篇论文提出了COT Fine-tuned框架，用于检测AI生成的文本并识别其背后的特定语言模型。该框架采用双任务方法：任务A区分AI生成与人类写作的文本，任务B识别生成文本的具体大语言模型(LLM)。其创新点在于引入思维链(Chain-of-Thought)推理，使模型能为其预测生成解释，提高了透明度和可解释性。实验表明，该方法在两项任务中都实现了高准确率，尤其在LLM识别和人机文本分类方面表现突出。研究证实思维链推理显著提升了模型的有效性和可解释性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-25T04:01:16Z
- **目录日期**: 2025-04-25
