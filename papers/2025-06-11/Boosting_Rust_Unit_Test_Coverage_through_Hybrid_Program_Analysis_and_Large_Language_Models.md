# Boosting Rust Unit Test Coverage through Hybrid Program Analysis and Large Language Models

**URL**: http://arxiv.org/abs/2506.09002v1

## 原始摘要

Unit testing is essential for ensuring software reliability and correctness.
Classic Search-Based Software Testing (SBST) methods and concolic
execution-based approaches for generating unit tests often fail to achieve high
coverage due to difficulties in handling complex program units, such as
branching conditions and external dependencies. Recent work has increasingly
utilized large language models (LLMs) to generate test cases, improving the
quality of test generation by providing better context and correcting errors in
the model's output. However, these methods rely on fixed prompts, resulting in
relatively low compilation success rates and coverage. This paper presents
PALM, an approach that leverages large language models (LLMs) to enhance the
generation of high-coverage unit tests. PALM performs program analysis to
identify branching conditions within functions, which are then combined into
path constraints. These constraints and relevant contextual information are
used to construct prompts that guide the LLMs in generating unit tests. We
implement the approach and evaluate it in 10 open-source Rust crates.
Experimental results show that within just two or three hours, PALM can
significantly improves test coverage compared to classic methods, with
increases in overall project coverage exceeding 50% in some instances and its
generated tests achieving an average coverage of 75.77%, comparable to human
effort (71.30%), highlighting the potential of LLMs in automated test
generation. We submitted 91 PALM-generated unit tests targeting new code. Of
these submissions, 80 were accepted, 5 were rejected, and 6 remain pending
review. The results demonstrate the effectiveness of integrating program
analysis with AI and open new avenues for future research in automated software
testing.


## AI 摘要

传统基于搜索的软件测试（SBST）和符号执行方法在生成高覆盖率单元测试时面临挑战，难以处理复杂分支条件和外部依赖。本文提出PALM方法，通过程序分析提取函数分支条件形成路径约束，结合上下文构建动态提示来指导大语言模型（LLM）生成测试用例。在10个Rust开源项目实验中，PALM仅用2-3小时就将项目覆盖率提升超50%，生成测试平均覆盖率达75.77%（接近人工的71.30%）。提交的91个测试中有80个被接受，验证了该方法在自动化测试中的有效性，为AI与程序分析的结合开辟了新方向。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-11T08:02:04Z
- **目录日期**: 2025-06-11
