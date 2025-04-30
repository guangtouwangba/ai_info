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

AegisLLM是一种多智能体协同防御系统，用于抵御对抗攻击和信息泄露。该系统通过四个自主智能体（协调器、偏转器、响应器和评估器）的协作流程，确保大语言模型输出的安全合规，并通过提示优化实现自我改进。研究表明，通过增加智能体角色和自动提示优化（如DSPy），该系统显著提升了鲁棒性且不影响模型效用。在WMDP遗忘基准测试中，仅需20个训练样本和不到300次调用即可实现近乎完美的遗忘；在越狱测试中，StrongReject性能提升51%，PHTest误拒率仅7.9%。代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-30T10:01:59Z
- **目录日期**: 2025-04-30
