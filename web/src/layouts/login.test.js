import { render, screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { toBeInTheDocument } from "@testing-library/jest-dom";
import LayoutLogin from "./login";
import { BrowserRouter, Routes, Route } from "react-router-dom";

describe("Login Layout", () => {
  test("render error message if Login fails", async () => {
    window.fetch = jest.fn();
    window.fetch.mockResolvedValueOnce({
      json: async () =>
        Promise.resolve({ message: "invalid email or password" }),
    });

    render(
      <BrowserRouter>
        <Routes>
          <Route path="*" element={<LayoutLogin />} />
        </Routes>
      </BrowserRouter>
    );

    const email = screen.getByLabelText("Email");
    await userEvent.type(email, "test@test.com");

    const password = screen.getByLabelText("Password");
    await userEvent.type(password, "12");

    await userEvent.click(screen.getByTestId("login-button"));

    expect(fetch).toHaveBeenCalledTimes(1);
    expect(
      await screen.findByText(/Incorrect user or Password/i)
    ).toBeInTheDocument();
  });
});
