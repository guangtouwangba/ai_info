# AegisLLM: Scaling Agentic Systems for Self-Reflective Defense in LLM Security

**URL**: http://arxiv.org/abs/2504.20965v1

## 原始摘要

We introduce AegisLLM, a cooperative multi-agent defense against adversarial
attacks and information leakage. In AegisLLM, a structured workflow of
autonomous agents - orchestrator, deflector, responder, and evaluator -
collaborate to ensure safe and compliant LLM outputs, while self-improving over
time through prompt optimization. We show that scaling agentic reasoning system
at test-time - both by incorporating additional agent roles and by leveraging
automated prompt optimization (such as DSPy)- substantially enhances robustness
without compromising model utility. This test-time defense enables real-time
adaptability to evolving attacks, without requiring model retraining.
Comprehensive evaluations across key threat scenarios, including unlearning and
jailbreaking, demonstrate the effectiveness of AegisLLM. On the WMDP unlearning
benchmark, AegisLLM achieves near-perfect unlearning with only 20 training
examples and fewer than 300 LM calls. For jailbreaking benchmarks, we achieve
51% improvement compared to the base model on StrongReject, with false refusal
rates of only 7.9% on PHTest compared to 18-55% for comparable methods. Our
results highlight the advantages of adaptive, agentic reasoning over static
defenses, establishing AegisLLM as a strong runtime alternative to traditional
approaches based on model modifications. Code is available at
https://github.com/zikuicai/aegisllm


## AI 摘要

AegisLLM是一种基于多智能体协作的防御系统，用于对抗对抗性攻击和信息泄露。该系统包含四个自主智能体（协调器、偏转器、响应器和评估器），通过结构化工作流程确保大语言模型输出的安全性和合规性，并利用提示优化实现自我改进。实验表明，AegisLLM在无需重新训练模型的情况下，显著提升了鲁棒性。在WMDP遗忘基准测试中，仅用20个训练样本和不到300次语言模型调用就实现了接近完美的遗忘效果；在越狱测试中，性能较基线模型提升51%，误拒率仅为7.9%。该方法展现了动态智能体防御优于静态防御的优势。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-30T14:01:53Z
- **目录日期**: 2025-04-30
